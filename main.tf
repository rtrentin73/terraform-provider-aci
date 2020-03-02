provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102/"
  insecure = true
}

# resource "aci_user" "demo1" {
#   name          = "user_demo1"
#   pwd           = "123456786"
#   description   = "This user is created by terraform"
# }

resource "aci_maintGrp" "demo1" {
  name          = "maintGrp1"
}

resource "aci_maintP" "demo1" {
  name = "maintP1"
}
