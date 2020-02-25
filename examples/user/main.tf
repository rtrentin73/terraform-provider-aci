 provider "aci" {
   username = "admin"
   password = "cisco123"
   private_key = ""
   cert_name = ""
   url      = ""
   insecure = true
}

resource "aci_user" "demouser" {
  name = "demo_user"
  description = "First user"
}
