package config

import "os"

var HalcyonMDB = os.Getenv("HALCYON_MDB")
var VersionColl = os.Getenv("VERSION_COLL")
var BrandColl = os.Getenv("BRAND_COLL")
var DeviceColl = os.Getenv("DEVICE_COLL")
var MaintainerColl = os.Getenv("MAINTAINER_COLL")
var ApplyFormColl = os.Getenv("MAINTAINER_COLL")
var AccountsColl = os.Getenv("ACCOUNTS_COLL")
var IssuesColl = os.Getenv("ISSUES_COLL")
var CommentsColl = os.Getenv("COMMENTS_COLL")
