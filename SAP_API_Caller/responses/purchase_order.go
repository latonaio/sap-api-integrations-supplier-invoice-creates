package responses

type PurchaseOrder struct {
	D struct {
			SupplierInvoice                string `json:"SupplierInvoice"`
			FiscalYear                     string `json:"FiscalYear"`
			SupplierInvoiceItem            string `json:"SupplierInvoiceItem"`
			PurchaseOrder                  string `json:"PurchaseOrder"`
			PurchaseOrderItem              string `json:"PurchaseOrderItem"`
			Plant                          string `json:"Plant"`
			TaxCode                        string `json:"TaxCode"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			SupplierInvoiceItemAmount      string `json:"SupplierInvoiceItemAmount"`
			PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
			QuantityInPurchaseOrderUnit    string `json:"QuantityInPurchaseOrderUnit"`
			PurchaseOrderPriceUnit         string `json:"PurchaseOrderPriceUnit"`
			QtyInPurchaseOrderPriceUnit    string `json:"QtyInPurchaseOrderPriceUnit"`
			SupplierInvoiceItemText        string `json:"SupplierInvoiceItemText"`
			PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
	} `json:"d"`
}

