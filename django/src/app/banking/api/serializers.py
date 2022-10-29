from rest_framework import serializers

from ..models import Client, Card, Contract, Credit, PaymentSchedule


class ClientSerializer(serializers.ModelSerializer):
    class Meta:
        model = Client
        fields = ['name', 'phone', 'registration_address', 'residential_address', 'client_type', 'inn', 'ogrn', 'kpp']


class CardSerializer(serializers.ModelSerializer):
    class Meta:
        model = Card
        fields = ['card_id', 'balance', 'valid_to', 'card_type', 'currency', 'client']


class ContractSerializer(serializers.ModelSerializer):
    class Meta:
        model = Contract
        fields = ['contract_type', 'conclusion_date', 'contract_content', 'client']


class CreditSerializer(serializers.ModelSerializer):
    class Meta:
        model = Credit
        fields = ['summ', 'percent', 'conclusion_date', 'end_date', 'contract']


class PaymentScheduleSerializer(serializers.ModelSerializer):
    class Meta:
        model = PaymentSchedule
        fields = [
            'total_sum',
            'currency',
            'commission',
            'repayment_of_interest_sum',
            'sum_repayment_loan_part',
            'date_begin',
            'date_end',
            'contract'
        ]
