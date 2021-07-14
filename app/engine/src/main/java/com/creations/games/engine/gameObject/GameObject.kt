package com.creations.games.engine.gameObject

import com.badlogic.gdx.scenes.scene2d.ui.WidgetGroup

/**
 * Base object
 */
abstract class GameObject: WidgetGroup() {
    open fun touchDown(x: Float, y: Float, pointer: Int, button: Int): Boolean = false
    open fun touchUp(x: Float, y: Float, pointer: Int, button: Int){}
    open fun touchDragged(x: Float, y: Float, pointer: Int){}

    //key calls requires setting keyboard focus to be set on the respective objects
    open fun keyDown(keycode: Int): Boolean = false
    open fun keyUp(keycode: Int): Boolean = false

    //right click
    open fun rightClicked(x: Float, y: Float){}
}