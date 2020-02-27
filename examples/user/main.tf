provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102/"
  insecure = true
}

resource "aci_user" "demo1" {
    name          = "user_demo"
    password      = "123456"
    description   = "This user is created by terraform"
}
