from rest_framework.decorators import api_view, authentication_classes, permission_classes
from rest_framework.response import Response

from ..models import User


@api_view(['POST'])
@permission_classes([])
@authentication_classes([])
def signup(request):
    User.objects.create_user(
        phone=request.data['phone'],
        email=request.data.get('email'),
        password=request.data['password']
    )
    return Response(status=201)
