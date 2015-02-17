### Demo code and slides presented at Gophercon India 2015

The demo showcases a typical EMBD based application. The setup has a Raspberry Pi 2 connected to a L3GD2 sensor which is used as a gyroscope to measure orientation. This data is then sent to a Qt based application using mangos (Go implementation of nanomsg). The application renders a image which follows the orientation of the sensor.