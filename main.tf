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



resource "aci_user" "user3" {
  name = "demouser5"
  pwd = "12345678"
  description="kjfas"
}




resource "aci_usercert_2" "cert2" {
  name = "certificate_2"
  description = "njkdcdsfdfdfd"
  local_user_dn = "${aci_user.user3.id}"
  file_path = "C:/Windows/System32/public.cer"
  
}


