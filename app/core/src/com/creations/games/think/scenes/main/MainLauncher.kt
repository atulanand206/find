package com.creations.games.think.scenes.main

import com.badlogic.gdx.Gdx
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.dependency.assetManager
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.GameScene
import com.creations.games.think.scenes.icons.FrameIcon
import com.creations.games.think.scenes.icons.PlayIcon
import com.creations.games.think.scenes.icons.SettingsIcon
import com.creations.games.think.scenes.main.objects.Header
import com.creations.games.think.scenes.main.objects.PuzzleFrame
import com.creations.games.think.scenes.main.objects.ThemeSwitch
import com.creations.games.think.utils.assets

class MainLauncher(private val di: DI) : GameScene(di) {

	private lateinit var header: GameObject
	private lateinit var puzzle: GameObject

	private lateinit var frameIcon: GameObject
	private lateinit var playIcon: GameObject
	private lateinit var settingsIcon: GameObject

	private lateinit var switch: GameObject

	init {
		load()
		di.assetManager
	}

	override fun load() {
		addBackground(di.assets.imgMain)
		header = Header(di, xTheme)
		header.color = viewColor(xTheme.specs)
		header.setPosition(0f, Values.VIRTUAL_HEIGHT * 0.95f)

		puzzle = PuzzleFrame(di)
		puzzle.setPosition(0f, 0f)

		switch = ThemeSwitch(di)
		switch.setPosition(Values.VIRTUAL_WIDTH * 0.10f, Values.VIRTUAL_HEIGHT * 0.08f)

		playIcon = addPlayIcon()
		frameIcon = addFrameIcon()
		settingsIcon = addSettingsIcon()

		Gdx.input.inputProcessor = stage

		addObjToScene(puzzle)
		addObjToScene(frameIcon)
		addObjToScene(switch)
		addObjToScene(playIcon)
		addObjToScene(settingsIcon)
		addObjToScene(header)

		stage.keyboardFocus = puzzle
	}

	override fun unload() {
		background.remove()
		header.remove()
		puzzle.remove()
		frameIcon.remove()
		playIcon.remove()
		switch.remove()
		settingsIcon.remove()
	}

	private fun addPlayIcon(): GameObject = di.assets.addIcon(PlayIcon(di),
			viewColor(xTheme.specs), Values.VIRTUAL_WIDTH * 0.5f, Values.VIRTUAL_HEIGHT * 0.35f,
			Values.VIRTUAL_WIDTH * 0.40f, Values.VIRTUAL_HEIGHT * 0.20f)

	private fun addFrameIcon() = di.assets.addIcon(FrameIcon(di),
			viewColor(xTheme.specs), Values.VIRTUAL_WIDTH * 0.5f, Values.VIRTUAL_HEIGHT * 0.10f,
			Values.VIRTUAL_WIDTH * 0.08f, Values.VIRTUAL_HEIGHT * 0.04f)

	private fun addSettingsIcon(): GameObject = di.assets.addIcon(SettingsIcon(di),
			viewColor(xTheme.specs), Values.VIRTUAL_WIDTH * 0.90f, Values.VIRTUAL_HEIGHT * 0.10f,
			Values.VIRTUAL_WIDTH * 0.06f, Values.VIRTUAL_HEIGHT * 0.03f)

	/**
	 * Act is called at start of every frame.
	 * Takes dt as a parameter
	 * Used to simulate the game world
	 *
	 * dt - delta time. The time elapsed since last frame
	 */
	override fun act(dt: Float) {

	}
}