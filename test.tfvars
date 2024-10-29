#------------------------------------------------------------------------------
# Common
#------------------------------------------------------------------------------
env                           = "uat"  #dev, test, preprod, prod
env_code                      = "ewu"
custom_tags = {
    "environment"      = "UAT"
    "region"           = "westeurope"
    "ci"               = "NSCP Reporting"
    "capability area"  = "AnalyticsAndReporting"
    "platform"         = "SCM"
    "business unit"	   = "Maersk_L&S"
}

sku_name = "P1v2"
scaledown = "false"