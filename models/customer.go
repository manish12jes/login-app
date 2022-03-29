package models

import(
  "database/sql"
  "time"
)

type Customer struct {
  ID          	 	uint
  Name         		string
  Email        		string
  InvoiceName       sql.NullString
  Phone		     	string
  VendorCode 		string
  Gst 				sql.NullString
  Address 			sql.NullString
  Type				string 
  CreatedAt    		time.Time
  UpdatedAt    		time.Time
}
