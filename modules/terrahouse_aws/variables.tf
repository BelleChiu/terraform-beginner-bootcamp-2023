variable "user_uuid" {
  description = "User UUID"
  type        = string

  validation {
    condition     = can(regex("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", var.user_uuid))
    error_message = "User UUID must be in the format of a UUID (e.g., 123e4567-e89b-12d3-a456-426655440000)"
  }
}

variable "bucket_name" {
  description = "AWS S3 Bucket Name"
  type        = string

  validation {
    condition     = can(regex("^[a-z0-9.-]{3,63}$", var.bucket_name))
    error_message = "Bucket name must be 3-63 characters long and can only contain lowercase letters, numbers, hyphens, and periods."
  }
}

variable "index_html_filepath" {
  description = "File path to the index.html file"
  type        = string

# https://developer.hashicorp.com/terraform/language/functions/can
  validation {
    condition     = can(file(var.index_html_filepath))
    error_message = "The specified file does not exist."
  }
}


variable "error_html_filepath" {
  description = "File path to the error.html file"
  type        = string

# https://developer.hashicorp.com/terraform/language/functions/can
  validation {
    condition     = can(file(var.error_html_filepath))
    error_message = "The specified file does not exist."
  }
}


variable "content_version" {
  description = "Content version (positive integer starting at 1)"
  type        = number

  validation {
    condition     = var.content_version > 0 && floor(var.content_version) == var.content_version
    error_message = "Content version must be a positive integer starting at 1."
  }
}
