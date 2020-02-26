provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}


# resource "aci_user" "user1" {
#   name        = "demouser2"
#   pwd = "12345678"
#   description = "This tenant is created by terraform"
# }

resource "aci_user" "user2" {
  name = "demouser4"
  pwd = "12345"
  description="kjfas"
}




resource "aci_UserCert_2" "cert2" {
  name = "certificate_2"
  description = "njkdcdsfdfdfd"
  user_name = "${aci_user.user2.id}"
}


