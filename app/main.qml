import QtQuick 2.4
import QtQuick.Window 2.2
import QtMultimedia 5.4
import QtGraphicalEffects 1.0

import com.sostronk.gophercon2015 0.1

Window {
    visible: true
    width: 800; height: 600

    Listener {
        id: l
    }

    Image {
        anchors.fill: parent
        source: "ground.jpg"

        Camera {
            id: camera
        }

        Item {
            anchors { top: parent.top; right: parent.right; margins: 10 }
            width: 300; height: width*(9/16)
            clip: true

            VideoOutput {
                id: video
                anchors.centerIn: parent
                width: parent.width*2; height: parent.height*2
                source: camera

                MouseArea {
                    anchors.fill: parent
                    onClicked: negativeEffect.visible = !negativeEffect.visible
                }
            }

            LevelAdjust {
                id: negativeEffect
                anchors.fill: video
                source: video
                minimumOutput: "#00ffffff"
                maximumOutput: "#ff000000"
            }
        }

        Image {
            id: rect

            anchors { left: parent.left; bottom: parent.bottom; margins: 10 }
            width: 500; height: 425

            source:"plane.png"
            smooth: true
            mipmap: true
            transform: [
                Rotation {
                    origin.y: rect.height/2; origin.x: rect.width/2; origin.z: 0
                    axis { x: 0; y:0 ; z: 1 }
                    angle: -l.z

                    Behavior on angle {
                        NumberAnimation {
                            duration: 200
                        }
                    }
                },
                Rotation {
                    origin.y: rect.height/2; origin.x: rect.width/2; origin.z: 0
                    axis { x: 0; y: 1; z: 0 }
                    angle: l.y

                    Behavior on angle {
                        NumberAnimation {
                            duration: 200
                        }
                    }
                }
            ]
        }
    }
}
