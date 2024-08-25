User {
  name = "John Doe"
  age = 30
  is_active = true
  repeated roles = "string"
  address = Address
  repeated contacts = Contact
}

Address {
  street = "string"
  city = "string"
  country = "string"
}

Contact {
  type = "string"
  value = "string"
}
