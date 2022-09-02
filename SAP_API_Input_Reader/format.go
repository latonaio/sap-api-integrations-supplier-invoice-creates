package sap_api_input_reader

import (
	"sap-api-integrations-supplier-invoice-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.SupplierInvoice
	return &requests.Header{
		SupplierInvoice:               data.SupplierInvoice,
		FiscalYear:                    data.FiscalYear,
		CompanyCode:                   data.CompanyCode,
		DocumentDate:                  data.DocumentDate,
		PostingDate:                   data.PostingDate,
		SupplierInvoiceIDByInvcgParty: data.SupplierInvoiceIDByInvcgParty,
		InvoicingParty:                data.InvoicingParty,
		DocumentCurrency:              data.DocumentCurrency,
		InvoiceGrossAmount:            data.InvoiceGrossAmount,
		DocumentHeaderText:            data.DocumentHeaderText,
		PaymentTerms:                  data.PaymentTerms,
		DueCalculationBaseDate:        data.DueCalculationBaseDate,
		NetPaymentDays:                data.NetPaymentDays,
		PaymentBlockingReason:         data.PaymentBlockingReason,
		AccountingDocumentType:        data.AccountingDocumentType,
		BPBankAccountInternalID:       data.BPBankAccountInternalID,
		SupplierInvoiceStatus:         data.SupplierInvoiceStatus,
		DirectQuotedExchangeRate:      data.DirectQuotedExchangeRate,
		SupplyingCountry:              data.SupplyingCountry,
		PaymentMethod:                 data.PaymentMethod,
		InvoiceReference:              data.InvoiceReference,
		SupplierPostingLineItemText:   data.SupplierPostingLineItemText,
		TaxIsCalculatedAutomatically:  data.TaxIsCalculatedAutomatically,
		BusinessArea:                  data.BusinessArea,
		SupplierInvoiceIsCreditMemo:   data.SupplierInvoiceIsCreditMemo,
		ReverseDocument:               data.ReverseDocument,
		ReverseDocumentFiscalYear:     data.ReverseDocumentFiscalYear,
	}
}

func (sdc *SDC) ConvertToPurchaseOrder() *requests.PurchaseOrder {
	dataSupplierInvoice := sdc.SupplierInvoice
	data := sdc.SupplierInvoice.PurchaseOrder
	return &requests.PurchaseOrder{
		SupplierInvoice:                dataSupplierInvoice.SupplierInvoice,
		FiscalYear:                     dataSupplierInvoice.FiscalYear,
		SupplierInvoiceItem:            data.SupplierInvoiceItem,
		PurchaseOrder:                  data.PurchaseOrder,
		PurchaseOrderItem:              data.PurchaseOrderItem,
		Plant:                          data.Plant,
		TaxCode:                        data.TaxCode,
		DocumentCurrency:               data.DocumentCurrency,
		SupplierInvoiceItemAmount:      data.SupplierInvoiceItemAmount,
		PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
		QuantityInPurchaseOrderUnit:    data.QuantityInPurchaseOrderUnit,
		PurchaseOrderPriceUnit:         data.PurchaseOrderPriceUnit,
		QtyInPurchaseOrderPriceUnit:    data.QtyInPurchaseOrderPriceUnit,
		SupplierInvoiceItemText:        data.SupplierInvoiceItemText,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
	}
}
