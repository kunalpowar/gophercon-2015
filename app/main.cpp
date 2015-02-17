#include <QGuiApplication>
#include <QQmlApplicationEngine>

#include <QtQml>

#include "listener.h"

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);

    qmlRegisterType<Listener>("com.sostronk.gophercon2015", 0, 1, "Listener");
    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:/main.qml")));

    return app.exec();
}
