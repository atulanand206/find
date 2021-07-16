package com.creations.games.think.scenes.main

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.scenes.Scene
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.scenes.Background
import com.creations.games.think.scenes.game.gameObjects.Ball
import com.creations.games.think.scenes.game.gameObjects.Ground
import com.creations.games.think.utils.assets

class MainLauncher(private val di: DI) : Scene(di){
	private val assets = di.assets

	private val groundHeight = 20f;
	private val ballRadius = 20f;

	//gameObjects
	private lateinit var ball: Ball
	private lateinit var ground: Ground

	//init is the constructor
	init {
		background(Background(di))
		addBall()
		addGround()
		addTitle()
	}

	private fun addBall() {
		//create a new ball instance
		ball = Ball(di)
		ball.setRadius(ballRadius)
		ball.setPosition(screenCenter.x, screenCenter.y, Align.center)

		//todo.note - uncomment below line to see ball's drawing bounds. You can use (ctrl + /) on a line to toggle comment
//        ball.debug = true

		addObjToScene(ball)
		addObjToScene(ball.counter)
	}

	private fun addGround() {
		ground = Ground(di)
		ground.setPosition(0f,0f)
		ground.setSize(Values.VIRTUAL_WIDTH, groundHeight)

		//todo.note - A gameObject will only be drawn if it is added to the scene
		addObjToScene(ground)
	}

	private fun addTitle() {
		//todo.note - you can add text using labels like below
		val title = assets.createLabel("Jumping Ball", size = FontSize.F14, color = Color.BLACK)

		//place label at top center of screen
		title.setPosition(Values.VIRTUAL_WIDTH/2f, Values.VIRTUAL_HEIGHT, Align.top)

//        label.debug = true


		addObjToScene(title)
	}

	/**
	 * Act is called at start of every frame.
	 * Takes dt as a parameter
	 * Used to simulate the game world
	 *
	 * dt - delta time. The time elapsed since last frame
	 */
	override fun act(dt: Float) {
		//check ball's collision with floor in a simple way
		if(ball.y < ground.y + groundHeight){
			ball.y = ground.y + groundHeight
			ball.bounce(false)
		}
		if (ball.x < 0f) {
			ball.x = 0f
			ball.bounce(true)
		}
		if (ball.x > Values.VIRTUAL_WIDTH) {
			ball.x = Values.VIRTUAL_WIDTH
			ball.bounce(true)
		}
	}
}