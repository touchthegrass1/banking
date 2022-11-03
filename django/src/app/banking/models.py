from django.contrib.auth.models import AbstractUser
from django.core.validators import MinValueValidator, MaxValueValidator
from django.db import models
from django.utils.timezone import now

from .managers import UserManager


class User(AbstractUser):
    objects = UserManager()
    username = None

    phone = models.CharField(max_length=12, unique=True)

    USERNAME_FIELD = 'phone'

    class Meta:
        db_table = 'user'
        verbose_name = 'User'
        verbose_name_plural = 'Users'


class Client(models.Model):
    client_id = models.BigAutoField(primary_key=True)
    registration_address = models.CharField(max_length=300)
    residential_address = models.CharField(max_length=300)
    client_type = models.CharField(max_length=20)
    ogrn = models.CharField(max_length=20, unique=True)
    inn = models.CharField(max_length=20, unique=True)
    kpp = models.CharField(max_length=20)
    user = models.OneToOneField('User', on_delete=models.RESTRICT)

    class Meta:
        managed = False
        db_table = 'client'
        verbose_name = 'Client'
        verbose_name_plural = 'Clients'


class Card(models.Model):
    class CardType(models.TextChoices):
        credit = 'credit'
        debit = 'debit'

    card_id = models.CharField(primary_key=True, max_length=12)
    balance = models.DecimalField(default=0, decimal_places=2, max_digits=10)
    valid_to = models.DateTimeField()
    cvc_code = models.IntegerField(validators=[MinValueValidator(0), MaxValueValidator(999)])
    card_type = models.CharField(choices=CardType.choices, default=CardType.debit, max_length=6)
    currency = models.CharField(max_length=3)
    client = models.ForeignKey(Client, on_delete=models.RESTRICT, related_name='clients')

    class Meta:
        managed = False
        db_table = 'card'
        verbose_name = 'Card'
        verbose_name_plural = 'Cards'


class Contract(models.Model):
    class ContractType(models.TextChoices):
        loan_agreement = 'loan_agreement'
        bank_account_agreement = 'bank_account_agreement'
        settlement_and_cash_service_agreement = 'settlement_and_cash_service_agreement'

    contract_id = models.BigIntegerField(primary_key=True)
    contract_type = models.CharField(
        choices=ContractType.choices,
        default=ContractType.settlement_and_cash_service_agreement,
        max_length=37
    )
    conclusion_date = models.DateField(default=now)
    contract_content = models.TextField()
    client = models.ForeignKey(Client, on_delete=models.RESTRICT, related_name='contracts')

    class Meta:
        managed = False
        db_table = 'contract'
        verbose_name = 'Contract'
        verbose_name_plural = 'Contracts'


class Credit(models.Model):
    credit_id = models.BigIntegerField(primary_key=True)
    summ = models.DecimalField(decimal_places=2, max_digits=10)
    percent = models.DecimalField(decimal_places=2, max_digits=2)
    conclusion_date = models.DateField(default=now)
    end_date = models.DateField()
    contract = models.OneToOneField(Contract, on_delete=models.RESTRICT)

    class Meta:
        managed = False
        db_table = 'credit'
        verbose_name = 'Credit'
        verbose_name_plural = 'Credits'


class PaymentSchedule(models.Model):
    payment_schedule_id = models.BigIntegerField(primary_key=True)
    total_sum = models.DecimalField(decimal_places=2, max_digits=10)
    currency = models.CharField(max_length=2)
    commission = models.DecimalField(decimal_places=2, max_digits=2)
    repayment_of_interest_sum = models.DecimalField(decimal_places=2, max_digits=10)
    sum_repayment_loan_part = models.DecimalField(decimal_places=2, max_digits=10)
    date_begin = models.DateField()
    date_end = models.DateField()
    contract = models.ForeignKey(Contract, on_delete=models.RESTRICT, related_name='payment_schedules')

    class Meta:
        managed = False
        db_table = 'payment_schedule'
        verbose_name = 'Payment Schedule'
        verbose_name_plural = 'Payment Schedules'


class Transaction(models.Model):
    transaction_id = models.BigIntegerField(primary_key=True)
    transaction_type = models.CharField(max_length=10)
    card_from = models.ForeignKey(Card, on_delete=models.DO_NOTHING, related_name='card_from_ids')
    card_to = models.ForeignKey(Card, on_delete=models.DO_NOTHING, related_name='card_to_ids')
    card = models.ForeignKey(Card, on_delete=models.DO_NOTHING, related_name='card_ids')
    summ = models.DecimalField(max_digits=10, decimal_places=2)
    transaction_datetime = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'transaction'
        verbose_name = 'Transaction'
        verbose_name_plural = 'Transactions'
