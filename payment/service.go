package payment

import (
	"crowdfunding/campaign"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	transactionRepository transaction.Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
	ProcessPayment(input transaction.TransactionNotificationInput) error
}

func NewService(transactionRepository transaction.Repository, campaignRepository campaign.Repository) *service {
	return &service{transactionRepository, campaignRepository}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-75fgSQZ2S0SPSVXWuzwIraMH"
	midclient.ClientKey = "SB-Mid-client-Q0C0mWmxeKWybmvr"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway {
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID), //karena OrderID string jadi harus di convert
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil  

}

func (s *service) ProcessPayment(input transaction.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.transactionRepository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if (input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept") {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expired" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancel"
	}

	updatedTranscation, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTranscation.CampaignID)
	if err != nil {
		return err 
	}

	if updatedTranscation.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTranscation.Amount 
	
		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil

}