package main

import (
    "crypto/sha256"
    "flag"
    "math/big"
	// "encoding/json"
    "fmt"
    // "time"
)

// This is the 1024-bit MODP group from RFC 5114, section 2.1:
const primeHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
const generatorHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"

func fromHex(hex string) *big.Int {
    n, ok := new(big.Int).SetString(hex, 16)
    if !ok {
        panic("failed to parse hex number")
    }
    return n
}

var (
    numT *int
    numN *int
    numG *big.Int
    numP *big.Int
    priv []*PrivateKey
    user *User
)

func init() {
    numT = flag.Int("t", 7, "t of (t, n).")
    numN = flag.Int("n", 10, "n of (t, n).")
    flag.Parse()
    numG = fromHex(generatorHex)
    numP = fromHex(primeHex)

    // p（P）阶循环群 G，g（G）为生成元，s_i（X）为私钥
    for i := 1; i <= *numN; i += 1 {
        priv = append(priv, &PrivateKey{
            PublicKey: PublicKey{
                G: fromHex(generatorHex),
                P: fromHex(primeHex),
            },
            X: big.NewInt(int64(i)),
            AtriXkey: big.NewInt(int64(i * 2)),
        })
    }

    for _, key := range priv {
        key.AtriYkey = new(big.Int).Exp(key.G, key.AtriXkey, key.P)
        key.Y = new(big.Int).Exp(key.G, key.X, key.P)
    }
    var SI []*big.Int
    var EI []*big.Int
    for i := 1; i <= *numN; i += 1 {
        EI = append(EI, big.NewInt(int64(i)))
        SI = append(SI, new(big.Int).Add(priv[i - 1].AtriXkey, new(big.Int).Mul(EI[i-1], priv[i - 1].X)))
    }

    user = &User{
        X: big.NewInt(7),
        Y: new(big.Int).Exp(numG, big.NewInt(7), numP),
        M: "Signature Info",
        S: SI,
        E: EI,
    }

}

func VCGenerate(){
    M := []byte(user.M)

    // (t, n) 门限
    var R []*big.Int
    var T []*big.Int
    // 属于属性的 1~t
    for i := 1; i <= *numT; i += 1 {
        T = append(T, big.NewInt(10 + int64(i)))
        // fmt.Println(T[i - 1])
    }
    for _, t := range T {
        R = append(R, new(big.Int).Exp(numG, t, numP))
    }

    // 不属于属性的 t+1～n
    var C []*big.Int
    var D []*big.Int
    for i := 1; i <= *numN - *numT; i += 1 {
        C = append(C, big.NewInt(30 + int64(i)))
        D = append(D, big.NewInt(40 + int64(i)))
    }
    for i, c := range C {
        yi :=  new(big.Int).Mul(priv[i + *numT].AtriYkey, new(big.Int).Exp(priv[i + *numT].Y, user.E[i + *numT], numP))
        R = append(R, new(big.Int).Mul(new(big.Int).Exp(numG, D[i], numP), new(big.Int).Exp(yi, c, numP)))
    }
    var Ti []*big.Int
    for _, r := range R {
        Ti = append(Ti, new(big.Int).Exp(r, user.X, numP))
    }
    var ri []*big.Int
    for i := 1; i <= *numN; i += 1{
        ri = append(ri, priv[i-1].AtriYkey)
    }
    buf := M
    for _, t := range Ti {
        buf = append(buf, t.Bytes()...)
    }
    buf = append(buf, user.Y.Bytes()...)
    result := sha256.Sum256(buf)
    resultTemp := result[:]
    // fmt.Printf("%x", result)

    lagPoints := []*LagPoint {
        {
            X: big.NewInt(0),
            Y: new(big.Int).SetBytes(resultTemp),
        },
    }
    for i := 1; i <= *numN - *numT; i += 1 {
        lagPoints = append(lagPoints, &LagPoint{
            X: big.NewInt(int64(*numT + i)),
            Y: C[i - 1],
        })
    }

    var CTemp []*big.Int
    var DTemp []*big.Int
    for i := 1; i <= *numT; i += 1 {
        cTemp := LagRange(lagPoints, big.NewInt(int64(i)))
        CTemp = append(CTemp, cTemp)
        dTemp := new(big.Int).Sub(T[i - 1], new(big.Int).Mul(cTemp, priv[i - 1].X))
        DTemp = append(DTemp, dTemp)
    }

    C = append(CTemp, C...)
    D = append(DTemp, D...)
    // for i, c := range C {
    //   c.Mod(c, numP)
    //   D[i] = new(big.Int).Mod(D[i], numP)
    //   R[i] = new(big.Int).Mod(R[i], numP)
    // }
    // fmt.Println(C)
    // fmt.Println(D)
    // fmt.Println(R)
    user.AbsSignature = &ABSSignature{
        C: C,
        D: D,
        R: R,
        T: Ti,
        LagPoints: lagPoints,
    }
    // bData, _ := json.Marshal(returnInfo)
    // fmt.Println(string(bData))
    zkc1 := new(big.Int).Exp(numG, big.NewInt(int64(11)), numP)
    var zkc2 []*big.Int
    var zkw1 []*big.Int
    var zkz1 []*big.Int
    for i := 1; i <= *numN; i += 1 {
        zkc2 = append(zkc2, new(big.Int).Exp(R[i-1], big.NewInt(int64(11)), numP))
        tempBytes := zkc1.Bytes()
        hash := sha256.Sum256(append(tempBytes, zkc2[i-1].Bytes()...))
        // fmt.Println(hash)
        // zkw1 = append(zkw1, fromHex(string(hash)))
        zkw1 = append(zkw1, new(big.Int).SetBytes(hash[:]))
        // fmt.Println(zkw1[i - 1])
        zkz1 = append(zkz1,  new(big.Int).Mod(new(big.Int).Add(big.NewInt(int64(11)),new(big.Int).Mul(zkw1[i-1], user.X)), numP))
    }
    // zkProof := &ZkProof{
    //     c1: zkc1,
    //     c2: zkc2,
    //     w1: zkw1,
    //     z1: zkz1,
    // }
    user.ZkProof = &ZkProof{
        C1: zkc1,
        C2: zkc2,
        W1: zkw1,
        Z1: zkz1,
    }
    user.Wlast = new(big.Int).Exp(numG, big.NewInt(int64(10)), numP)
    user.Cnt = big.NewInt(int64(1))
    return
}

// func Verify() bool {

// }

func Verify() bool {
    C := user.AbsSignature.C
    // D := user.AbsSignature.D
    R := user.AbsSignature.R
    T := user.AbsSignature.T
    lagPoints := user.AbsSignature.LagPoints

    // 验证
    for i, cT := range C {
        res := cT.Cmp(LagRange(lagPoints, big.NewInt(int64(i+1))))
        if res != 0 {
            fmt.Println("ci failed;")
            return false
        }
    }
    buf := []byte(user.M)
    for _, t := range T {
        buf = append(buf, t.Bytes()...)
    }
    buf = append(buf, user.Y.Bytes()...)
    result := sha256.Sum256(buf)
    hashInt := new(big.Int).SetBytes(result[:])
    res := hashInt.Cmp(LagRange(lagPoints, big.NewInt(0)))
    if res != 0 {
        fmt.Println("Hash failed;")
        return false
    }
    // fmt.Println("Signature success;")
    
    zkc1 := user.ZkProof.C1
    zkc2 := user.ZkProof.C2
    zkw1 := user.ZkProof.W1
    zkz1 := user.ZkProof.Z1
    tempBytes := zkc1.Bytes()
    for i := 1; i <= *numN; i += 1 {
        hash := sha256.Sum256(append(tempBytes, zkc2[i-1].Bytes()...))
        // fmt.Println(hash)
        // zkw1 = append(zkw1, fromHex(string(hash)))
        res := zkw1[i - 1].Cmp(new(big.Int).SetBytes(hash[:]))
        if res != 0 {
            fmt.Println("Zk hash failed;")
            return false
        }
        
        g_Z := new(big.Int).Exp(numG, zkz1[i - 1], numP)
        res = g_Z.Cmp(new(big.Int).Mod(new(big.Int).Mul(zkc1, new(big.Int).Exp(user.Y, zkw1[i - 1], numP)), numP))
        if res != 0 {
            fmt.Println("Zk1 compare failed;")
            return false
        }
        R_Z := new(big.Int).Exp(R[i-1], zkz1[i-1], numP)
        res = R_Z.Cmp(new(big.Int).Mod(new(big.Int).Mul(zkc2[i-1], new(big.Int).Exp(T[i - 1], zkw1[i - 1], numP)), numP))
        if res != 0 {
            fmt.Println("Zk2 compare failed;")
            return false
        }
    }
    // fmt.Println("Zkproof success;")
    // fmt.Println("Cntproof success;")
    return true
}

func VCExtend(){
    M := []byte(user.M)
    *numT = *numT + 1
    user.AbsSignature.R[*numT - 1] = new(big.Int).Exp(numG, big.NewInt(10 + int64(*numT)), numP)
    user.AbsSignature.T[*numT - 1] = new(big.Int).Exp(user.AbsSignature.R[*numT - 1], user.X, numP)
    T := user.AbsSignature.R
    buf := M
    for _, t := range T {
        buf = append(buf, t.Bytes()...)
    }
    buf = append(buf, user.Y.Bytes()...)
    result := sha256.Sum256(buf)
    resultTemp := result[:]

    user.AbsSignature.LagPoints[0].Y = new(big.Int).SetBytes(resultTemp)
    lagPoints := append(user.AbsSignature.LagPoints[:1],user.AbsSignature.LagPoints[2:]...)
    user.AbsSignature.LagPoints = lagPoints
    for i := 1; i <= *numT; i += 1 {
        user.AbsSignature.C[i - 1] = LagRange(lagPoints, big.NewInt(int64(i)))
        user.AbsSignature.D[i - 1] = new(big.Int).Sub(T[i - 1], new(big.Int).Mul(user.AbsSignature.C[i- 1], priv[i - 1].X))
    }
    return
}

func main() {
    // sBegin := time.Now().UnixNano()
    VCGenerate()
    // sEnd := time.Now().UnixNano()
    // fmt.Println("VC生成时间")
    // fmt.Println(float64(sEnd-sBegin) / 1e9)
    // sign := Generate("Signature Info")
    // bData, _ := json.Marshal(user)
    // fmt.Println(string(bData))
    // sBegin = time.Now().UnixNano()
    // bi := new(big.Int).Mod(new(big.Int).Add(user.Wlast, big.NewInt(int64(11))), numP)
    // w := new(big.Int).Exp(numG, bi, numP)
    // cnt_m := new(big.Int).Mod(new(big.Int).Mul(user.Cnt, new(big.Int).Exp(priv[0].Y, bi, numP)), numP)
    // // 构建CntProof
    // user.CntProof = &CntProof{
    //     Wlast: user.Wlast,
    //     Cnt_m: cnt_m,
    //     W: w,
    // }
    // sEnd = time.Now().UnixNano()
    // fmt.Println("CNTProof生成时间")
    // fmt.Println(float64(sEnd-sBegin) / 1e9)
    
    // sBegin = time.Now().UnixNano()
    // Verify()
    // sEnd = time.Now().UnixNano()
    // fmt.Println("验证时间")
    // fmt.Println(float64(sEnd-sBegin) / 1e9)

    // sBegin = time.Now().UnixNano()
    // VCExtend()
    // sEnd = time.Now().UnixNano()
    // fmt.Println("VC扩充时间")
    // fmt.Println(float64(sEnd-sBegin) / 1e9)

    // var bData []byte
    // bData, _ = json.Marshal(user.AbsSignature)
    // fmt.Println("Signature大小")
    // fmt.Println(len(bData))
    // bData, _ = json.Marshal(user.ZkProof)
    // fmt.Println("ZKProof大小")
    // fmt.Println(len(bData))
    // bData, _ = json.Marshal(user.CntProof)
    // fmt.Println("CNTProof大小")
    // fmt.Println(len(bData))
}
