from django.contrib import admin

from .models import User, Client, Card, Contract, Credit, PaymentSchedule, Transaction


class ReadOnlyAdmin(admin.ModelAdmin):
    readonly_fields = []

    def get_readonly_fields(self, request, obj=None):
        return list(self.readonly_fields) + \
               [field.name for field in obj._meta.fields] + \
               [field.name for field in obj._meta.many_to_many]

    def has_add_permission(self, request):
        return False

    def has_delete_permission(self, request, obj=None):
        return False


admin.site.register(User)
admin.site.register(Client, ReadOnlyAdmin)
admin.site.register(Contract, ReadOnlyAdmin)
admin.site.register(Credit, ReadOnlyAdmin)
admin.site.register(PaymentSchedule, ReadOnlyAdmin)
admin.site.register(Transaction, ReadOnlyAdmin)
