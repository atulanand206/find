package com.creations.games.think.scenes.main

import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.data.Puzzle
import com.creations.games.think.data.Theme
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.Background
import com.creations.games.think.scenes.GameScene
import com.creations.games.think.scenes.icons.*
import com.creations.games.think.utils.assets

class MainLauncher(private val di: DI) : GameScene(di) {
	private val assets = di.assets

	private val themes = arrayListOf(xTheme)
	private val level = 0

	init {
		addObjToScene(Background(di))
		addBackIcon()
		addObjToScene(addTitle(xTheme))
		addPuzzle(xTheme.puzzles[level])
		addLevel(xTheme.puzzles[level])
		addPlayIcon()
		addFrameIcon()
		addLightIcon()
		addSettingsIcon()
	}

	private fun addPuzzle(puzzle: Puzzle) {
		addObjToScene(addPuzzleName(puzzle))
		addObjToScene(addPuzzleText(puzzle))
	}

	private fun addLevel(puzzle: Puzzle) {
		addObjToScene(addLevelNo(puzzle))
		addObjToScene(addLeftSwipeButton(puzzle))
		addObjToScene(addRightSwipeButton(puzzle))
	}

	private fun addBackIcon() {
		addObjToScene(di.assets.addIcon(BackIcon(di), viewColor(xTheme.specs),
				Values.VIRTUAL_WIDTH * 0.07f, Values.VIRTUAL_HEIGHT * 0.95f,
				Values.VIRTUAL_HEIGHT * 0.04f, Values.VIRTUAL_HEIGHT * 0.04f))
	}

	private fun addPlayIcon() {
		addObjToScene(di.assets.addIcon(PlayIcon(di), viewColor(xTheme.specs),
				Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.35f,
				Values.VIRTUAL_HEIGHT * 0.20f, Values.VIRTUAL_HEIGHT * 0.20f))
	}

	private fun addFrameIcon() {
		addObjToScene(di.assets.addIcon(FrameIcon(di), viewColor(xTheme.specs),
				Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.10f,
				Values.VIRTUAL_HEIGHT * 0.04f, Values.VIRTUAL_HEIGHT * 0.04f))
	}

	private fun addLightIcon() {
		addObjToScene(di.assets.addIcon(DarkIcon(di), viewColor(xTheme.specs),
				Values.VIRTUAL_WIDTH * 0.10f, Values.VIRTUAL_HEIGHT * 0.10f,
				Values.VIRTUAL_HEIGHT * 0.03f, Values.VIRTUAL_HEIGHT * 0.03f))
	}

	private fun addSettingsIcon() {
		addObjToScene(di.assets.addIcon(SettingsIcon(di), viewColor(xTheme.specs),
				Values.VIRTUAL_WIDTH * 0.90f, Values.VIRTUAL_HEIGHT * 0.10f,
				Values.VIRTUAL_HEIGHT * 0.03f, Values.VIRTUAL_HEIGHT * 0.03f))
	}

	private fun addTitle(theme: Theme): GameObject = assets.addLabel(
			theme.name, FontSize.F20, viewColor(theme.specs),
			Values.VIRTUAL_WIDTH * 0.95f, Values.VIRTUAL_HEIGHT * 0.95f, Align.right)

	private fun addPuzzleName(puzzle: Puzzle): GameObject = assets.addWrappedLabel(
			puzzle.name, FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.85f, Align.center)

	private fun addPuzzleText(puzzle: Puzzle): GameObject = assets.addWrappedLabel(
			puzzle.text, FontSize.F18, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.72f, Align.center)

	private fun addLevelNo(puzzle: Puzzle): GameObject = assets.addLabel(
			(level + 1).toString(), FontSize.F100, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.56f, Align.center)

	private fun addLeftSwipeButton(puzzle: Puzzle): GameObject = assets.addLabel(
			"{", FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH * 0.07f, Values.VIRTUAL_HEIGHT * 0.56f, Align.center)

	private fun addRightSwipeButton(puzzle: Puzzle): GameObject = assets.addLabel(
			"}", FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH * 0.93f, Values.VIRTUAL_HEIGHT * 0.56f, Align.center)

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