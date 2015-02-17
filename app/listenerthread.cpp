#include "listenerthread.h"

#include <QTimer>
#include <QDebug>

#include <iostream>
#include <system_error>
#include <nnxx/message.h>
#include <nnxx/pair.h>
#include <nnxx/socket.h>

ListenerThread::ListenerThread(QObject *parent)
{
    moveToThread(this);
}

ListenerThread::~ListenerThread()
{
    quit();
}

void ListenerThread::run()
{
    QTimer::singleShot(0, this, SLOT(receiveMessage()));
    exec();
}

void ListenerThread::receiveMessage()
{
    try {
        nnxx::socket s { nnxx::SP, nnxx::PAIR };
        const char *addr = "tcp://192.168.2.1:40899";

        s.bind(addr);
        s.send("START");
        nnxx::message msg = s.recv();

        auto dataStr = QString::fromStdString(to_string(msg));
        emit newData(dataStr);
    } catch (const std::system_error &e) {
        std::cerr << e.what() << std::endl;
    }

    QTimer::singleShot(0, this, SLOT(receiveMessage()));
}
