#include "listener.h"
#include "listenerthread.h"

#include <QTimer>
#include <QDebug>
#include <QStringList>

Listener::Listener(QObject *parent) : QObject(parent)
{
    auto l = new ListenerThread(this);
    connect(l, &ListenerThread::newData, [=](QString dataStr) {
        QStringList data = dataStr.split(",");
        m_x = data.at(0).toInt();
        m_y = data.at(1).toInt();
        m_z = data.at(2).toInt();
        emit xChanged(m_x);
        emit yChanged(m_y);
        emit zChanged(m_z);
    });
    l->start();
}

int Listener::x() const
{
    return m_x;
}

int Listener::y() const
{
    return m_y;
}

int Listener::z() const
{
    return m_z;
}

