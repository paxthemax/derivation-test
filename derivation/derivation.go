package derivation

// KeyDerivation derives a new private key
func KeyDerivation(pub *Key, priv *Key) Key {
	var point ExtendedGroupElement
	var point2 ProjectiveGroupElement
	var point3 CompletedGroupElement

	if !Sc_check(priv) {
		panic("Invalid private key.")
	}
	tmp := *pub
	if !point.FromBytes(&tmp) {
		panic("Invalid public key.")
	}

	tmp = *priv
	GeScalarMult(&point2, &tmp, &point)
	GeMul8(&point3, &point2)
	point3.ToProjective(&point2)

	point2.ToBytes(&tmp)
	return tmp
}
