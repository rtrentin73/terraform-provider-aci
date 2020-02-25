provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}


resource "aci_user" "user1" {
  name        = "demouser"
  pwd = "123456789"
  description = "This tenant is created by terraform"
}

