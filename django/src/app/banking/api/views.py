from django.db import transaction
from rest_framework.decorators import api_view, authentication_classes, permission_classes
from rest_framework.response import Response

from .serializers import UserCreateSerializer


@api_view(['POST'])
@permission_classes([])
@authentication_classes([])
def signup(request):
    with transaction.atomic():
        serializer = UserCreateSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(status=201)
        return Response(status=400, data=serializer.errors)
