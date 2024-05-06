package service

import (
	"errors"
	"fmt"
	"log"
	"otp-service/internal/gateway"
	v1 "otp-service/internal/gen/v1"
	"regexp"
	"strconv"
)

type CountryCode int32

const (
	India CountryCode = 91
)

var AllowedCountryCodes []CountryCode

func init() {
	AllowedCountryCodes = append(AllowedCountryCodes, India)
}

type IOtpService interface {
	SendMessage(request *v1.GenerateOTPRequest) error
}

func NewService(sender gateway.OTPSender, generator IGenerator) IOtpService {
	return &otpService{sender: sender, generator: generator}
}

type otpService struct {
	sender    gateway.OTPSender
	generator IGenerator
}

func (s *otpService) SendMessage(request *v1.GenerateOTPRequest) error {
	fmt.Printf("sending sms to %s\n", request.PhoneNumber)
	err := validateRequest(request)
	if err != nil {
		log.Println(err)
		return err
	}
	generatedOtp, err := s.generator.Generate(request.PhoneNumber)
	if err != nil {
		log.Println(err)
		return err
	}
	formattedNumber := "+" + strconv.Itoa(int(request.CountryCode)) + request.PhoneNumber
	err = s.sender.SendOTP(formattedNumber, generatedOtp)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func validateRequest(req *v1.GenerateOTPRequest) error {
	phErr := validatePhoneNumber(req.PhoneNumber)
	cErr := validateCountryCodes(req.CountryCode)
	return errors.Join(phErr, cErr)
}

func validatePhoneNumber(phoneNumber string) error {
	pattern := `^\+?\d{0,3}?\d{10}$`
	matched, _ := regexp.MatchString(pattern, phoneNumber)
	if !matched {
		return fmt.Errorf("phone number %s is not a valid number", phoneNumber)
	}
	return nil
}

func validateCountryCodes(countryCode int32) error {
	for _, code := range AllowedCountryCodes {
		if countryCode == int32(code) {
			return nil
		}
	}
	return fmt.Errorf("country code %d is not yet supported", countryCode)
}
