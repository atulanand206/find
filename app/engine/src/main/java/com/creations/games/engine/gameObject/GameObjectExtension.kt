package com.creations.games.engine.gameObject

import com.badlogic.gdx.Input
import com.badlogic.gdx.scenes.scene2d.InputEvent
import com.badlogic.gdx.scenes.scene2d.InputListener
import com.badlogic.gdx.scenes.scene2d.utils.ClickListener

fun GameObject.addInputListener() {
    val listener = object : InputListener() {
        override fun touchDown(event: InputEvent?, x: Float, y: Float, pointer: Int, button: Int): Boolean = touchDown(x, y, pointer, button)
        override fun touchUp(event: InputEvent?, x: Float, y: Float, pointer: Int, button: Int) = touchUp(x, y, pointer, button)
        override fun touchDragged(event: InputEvent?, x: Float, y: Float, pointer: Int) = touchDragged(x, y, pointer)
        override fun keyDown(event: InputEvent?, keycode: Int): Boolean = keyDown(keycode)
        override fun keyUp(event: InputEvent?, keycode: Int): Boolean = keyUp(keycode)
    }

    val clickListener = object : ClickListener(Input.Buttons.RIGHT){
        override fun clicked(event: InputEvent?, x: Float, y: Float) = rightClicked(x,y)
    }

    this.addListener(listener)
    this.addListener(clickListener)
}