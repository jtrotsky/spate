/*
{
  "id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
  "name": "Main Register",
  "outlet_id": "6c8f04b3-3110-11e3-a29a-bc305bf5da20",
  "print_receipt": "0",
  "email_receipt": "0",
  "ask_for_note_on_save": "1",
  "print_note_on_receipt": "0",
  "ask_for_user_on_sale": "0",
  "show_discounts_on_receipt": "1",
  "receipt_header": "<h1>piotr<h1>",
  "receipt_barcoded": "1",
  "receipt_footer": "<h1>Thanks for stopping by<h1>",
  "receipt_style_class": "has-receipt-80",
  "invoice_prefix": "",
  "invoice_suffix": "",
  "invoice_sequence": 50,
  "register_open_count_sequence": "4",
  "register_open_time": "2013-09-23 23:30:42",
  "register_close_time": "",
  "quick_keys_template": {},
  "receipt": {
    "fields": {
      "label_invoice": "Invoice #:",
      "label_invoice_title": "Receipt / Tax Invoice",
      "label_served_by": "Served by:",
      "label_line_discount": "Less discount ",
      "label_sub_total": "Subtotal",
      "label_tax": "Tax",
      "label_to_pay": "TO PAY",
      "label_total": "TOTAL",
      "label_change": "Change",
      "header": "<h1>piotr<h1>",
      "footer": "<h1>Thanks for stopping by<h1>"
    }
  }
}
*/

// Package types ...
package types

// Register ...
type Register struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	OutletID string `json:"outlet_id,omitempty"`
}
