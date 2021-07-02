"""paper_sever URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path
from learn import views as learn_views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('addConference', learn_views.addConference),
    path('addJournal', learn_views.addJournal),
    path('addSoftware', learn_views.addSoftware),
    path('addPatent', learn_views.addPatent),
    path('traceBackwardPaper', learn_views.traceBackwardPaper),
    path('traceBackwardConOrPatent', learn_views.traceBackwardConOrPatent),
    path('traceBackwardConOrPatentFromTxid', learn_views.traceBackwardConOrPatentFromTxid),
    path('traceBackwardPaperFromTxid', learn_views.traceBackwardPaperFromTxid),
    path('traceBackwardAllPaperFromHashTitle', learn_views.traceBackwardAllPaperFromHashTitle),
    path('traceBackwardAllConOrPatentFromHashTitle', learn_views.traceBackwardAllConOrPatentFromHashTitle),
    path('GetPaper', learn_views.GetPaper)
    # path('', learn_views.home, name="home"), 
]
