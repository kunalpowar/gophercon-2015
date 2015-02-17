TEMPLATE = app

QT += qml quick

SOURCES += main.cpp \
    listener.cpp \
    listenerthread.cpp

RESOURCES += qml.qrc

# Additional import path used to resolve QML modules in Qt Creator's code model
QML_IMPORT_PATH =

# Default rules for deployment.
include(deployment.pri)

CONFIG += c++11

macx: LIBS += -lnnxx

HEADERS += \
    listener.h \
    listenerthread.h
