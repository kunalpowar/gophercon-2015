#ifndef LISTENERTHREAD_H
#define LISTENERTHREAD_H

#include <QThread>

class ListenerThread : public QThread
{
    Q_OBJECT
public:
    explicit ListenerThread(QObject *parent = 0);
    ~ListenerThread();

signals:
    void newData(QString data);

    // QThread interface
protected:
    virtual void run();

public slots:
    void receiveMessage();
};

#endif // LISTENERTHREAD_H
