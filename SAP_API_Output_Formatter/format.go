package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-supplier-invoice-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
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

	return header, nil
}

func ConvertToPurchaseOrder(raw []byte, l *logger.Logger) (*PurchaseOrder, error) {
	pm := &responses.PurchaseOrder{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PurchaseOrder. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	purchaseOrder := &PurchaseOrder{
		SupplierInvoice:                data.SupplierInvoice,
		FiscalYear:                     data.FiscalYear,
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

	return purchaseOrder, nil
}
