import logging

import django.db
from django.db import transaction
from rest_framework.decorators import api_view, authentication_classes, permission_classes
from rest_framework.response import Response

from .serializers import UserCreateSerializer

logger = logging.getLogger(__name__)


@api_view(['POST'])
@permission_classes([])
@authentication_classes([])
def signup(request):
    attempts = 0
    error = None
    while attempts < 3:
        try:
            with transaction.atomic():
                serializer = UserCreateSerializer(data=request.data)
                if serializer.is_valid():
                    serializer.save()
                    return Response(status=201)
                return Response(status=400, data=serializer.errors)
        except django.db.Error as e:
            if e.__cause__.pgcode in ('40000', '40001', '40002', '40P01'):
                attempts += 1
                error = e
            else:
                logger.error(e)
                break
    logger.error(error)
    return Response(status=500)
