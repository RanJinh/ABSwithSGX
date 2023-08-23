package main

import "math/big"

type ZkProof struct{
    C1 *big.Int          `json:"c1"`
    C2 []*big.Int          `json:"c2"`
    W1  []*big.Int          `json:"w"`
    Z1  []*big.Int          `json:"z"`
}

type CntProof struct{
    Wlast *big.Int `json:"wlast"`
    Cnt_m *big.Int `json:"cnt_m"`
    W *big.Int `json:"w"`
}

type User struct{
    X               *big.Int        `json:"sk"`
    Y               *big.Int        `json:"pk"`
    M               string          `json:"m"`
    AbsSignature    *ABSSignature   `json:"ABSSignature"`
    ZkProof         *ZkProof        `json:"zkproof"`
    CntProof        *CntProof       `json:"cntproof"`
    S               []*big.Int      `json:"s"`
    E               []*big.Int      `json:"e"`
    Cnt             *big.Int        `json:"cnt"`
    Wlast           *big.Int        `json:"wlast"`
}

type LagPoint struct {
    X   *big.Int    `json:"x"`
    Y   *big.Int    `json:"y"`
}

type PublicKey struct {
    G, P, Y, AtriYkey *big.Int
}

type PrivateKey struct {
    PublicKey
    X *big.Int
    AtriXkey *big.Int
}

type ABSSignature struct {
    C           []*big.Int          `json:"c"`
    D           []*big.Int          `json:"d"`
    R           []*big.Int          `json:"r"`
    T           []*big.Int          `json:"t"`
    LagPoints   []*LagPoint         `json:"lagpoints"`
}