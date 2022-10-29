from .base import *

import psycopg2
from django_replicated.settings import *

ALLOWED_HOSTS = os.environ['DJANGO_ALLOWED_HOSTS'].split(',')

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': os.environ['POSTGRESQL_DB_NAME'],
        'USER': os.environ['POSTGRESQL_DB_USER'],
        'PASSWORD': os.environ['POSTGRESQL_DB_PASSWORD'],
        'HOST': os.environ['POSTGRESQL_MASTER_DB_HOST'],
        'PORT': os.environ['POSTGRESQL_DB_PORT'],

        'OPTIONS': {
            'isolation_level': psycopg2.extensions.ISOLATION_LEVEL_SERIALIZABLE
        },
    },
    'slave1': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': os.environ['POSTGRESQL_DB_NAME'],
        'USER': os.environ['POSTGRESQL_DB_USER'],
        'PASSWORD': os.environ['POSTGRESQL_DB_PASSWORD'],
        'HOST': os.environ['POSTGRESQL_REPLICA1_DB_HOST'],
        'PORT': os.environ['POSTGRESQL_DB_PORT'],

        'OPTIONS': {
            'isolation_level': psycopg2.extensions.ISOLATION_LEVEL_SERIALIZABLE
        },
    },
    'slave2': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': os.environ['POSTGRESQL_DB_NAME'],
        'USER': os.environ['POSTGRESQL_DB_USER'],
        'PASSWORD': os.environ['POSTGRESQL_DB_PASSWORD'],
        'HOST': os.environ['POSTGRESQL_REPLICA2_DB_HOST'],
        'PORT': os.environ['POSTGRESQL_DB_PORT'],

        'OPTIONS': {
            'isolation_level': psycopg2.extensions.ISOLATION_LEVEL_SERIALIZABLE
        },
    }
}

REPLICATED_DATABASE_SLAVES = ['slave1', 'slave2']
DATABASE_ROUTERS = ['django_replicated.router.ReplicationRouter']

CSRF_COOKIE_SECURE = os.environ['DJANGO_CSRF_COOKIE_SECURE'] == 'True'

CORS_ALLOW_ALL_ORIGINS = os.environ['DJANGO_CORS_ALLOW_ALL_ORIGINS'] == 'True'

SESSION_COOKIE_SECURE = os.environ['DJANGO_SESSION_COOKIE_SECURE'] == 'True'
