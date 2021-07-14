package com.creations.games.think.scenes.game.gameObjects

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.g2d.Batch
import com.badlogic.gdx.math.Vector2
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.think.asset.FontSize
import com.creations.games.think.utils.assets
import java.util.*

class Ball(private val di: DI): GameObject() {
    //todo.note - references to all textures are stored in di.assets
    //fetch a reference to a circle texture for my ball object
    private val tex = di.assets.circle
    //used to show the ball's bounce count
    val counter = di.assets.createLabel("0", size = FontSize.F14, color = Color.WHITE)

    private var elapsedTime = 0f
    private var isBlue = true
    private var bounces = 0f
    private var bounceTime = Date()
    private var grad = 0.1f
    private var gradIncrement = 0.01f

    //motion params
    //initial horizontal velocity
    private val velocity = Vector2(64f, 0f)
    //todo.note - gravity has a negative value as down is -ve in y axis. Unit is pixel/s^2
    val gravity = -200f

    init{
        //set color to blue
        color.set(0f,0f,1f,1f)
        isBlue = true
    }

    /**
     * Increments the bounce count if sufficient time between the contacts.
     *
     * horizontal - True if the bounce is from the walls in the horizontal direction
     */
    fun bounce(horizontal: Boolean) {
        if (Date().time - bounceTime.time > 100) {
            if (horizontal) velocity.x *= -1
            else velocity.y *= -1
            bounces += 1
            counter.setText(bounces.toInt().toString())
        }
        bounceTime = Date()
    }

    fun setRadius(rad: Float) {
        //todo.note - all gameObjects have width and height, for a ball they are equal to 2 times radius
        setSize(2*rad, 2*rad)
    }

    /**
     * Act is called at start of every frame.
     * Takes dt as a parameter
     * Used to simulate the game world
     *
     * dt - delta time. The time elapsed since last frame
     */
    override fun act(dt: Float) {
        //todo.note - up is +ve on y axis
        //update ball velocity based on total acceleration
        //since only gravity is present, only y velocity is updated
        velocity.y += gravity * dt

        //noticed unusual things by switching the order of these lines
        counter.setPosition(x + width / 3, y + height / 3)
        //move ball based on velocity
        this.moveBy(velocity.x * dt,velocity.y * dt)

        //slowly change the ball color
        if (grad <= 0f || grad >= 1f) gradIncrement *= -1
        grad += gradIncrement
        color.set(1f-grad,0f,0f+grad,1f)

    }

    /**
     * Draw is called after act each frame
     * Takes a batch for drawing textures to screen
     * can be used to refresh the screen by drawing objects to their new location calculated in act
     */
    override fun draw(batch: Batch, parentAlpha: Float) {
        //change the batch color to balls color.
        //todo.note - Batch color applies a tint to entire texture while drawing. circle texture is white so it is drawn in the selected color
        batch.color = color

        //draw circle tex at ball's position and size
        batch.draw(tex, x, y, width, height)

        //ALWAYS revert batch color back to white
        batch.color = Color.WHITE
    }
}