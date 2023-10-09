# variable "bucket_name" {
#     type        = string
# }

variable "hometown" {
  type = object ({
    public_path = string
    content_version=number
  })
}

variable "cookandbake" {
  type = object({
    public_path = string
    content_version=number
  })
}

variable "terratowns_endpoint" {
  type = string
}

variable "terratowns_access_token" {
  type    = string
}

variable "teacherseat_user_uuid" {
  type = string
}