from django.urls import path

from .api import views

app_name = 'banking'

urlpatterns = [
    path('signup/', views.signup, name='signup'),
]
