#ifndef LISTENER_H
#define LISTENER_H

#include <QObject>

class Listener : public QObject
{
    Q_OBJECT
    Q_PROPERTY(int x READ x NOTIFY xChanged)
    Q_PROPERTY(int y READ y NOTIFY yChanged)
    Q_PROPERTY(int z READ z NOTIFY zChanged)

    int m_x = 0;
    int m_z = 0;
    int m_y = 0;

public:
    explicit Listener(QObject *parent = 0);

    int x() const;
    int y() const;
    int z() const;

signals:
    // QThread interface
    void xChanged(int arg);
    void yChanged(int arg);
    void zChanged(int arg);
};

#endif // LISTENER_H
