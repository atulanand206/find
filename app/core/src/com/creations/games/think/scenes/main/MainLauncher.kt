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
import com.creations.games.think.utils.assets

class MainLauncher(private val di: DI) : GameScene(di) {
	private val assets = di.assets

	private val themes = arrayListOf(xTheme)
	private val level = 0

	init {
		addObjToScene(Background(di))
		addObjToScene(addTitle(xTheme))
		addPuzzle(xTheme.puzzles[level])
		addLevel(xTheme.puzzles[level])
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

	private fun addTitle(theme: Theme): GameObject = assets.addLabel(
			theme.name, FontSize.F20, viewColor(theme.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.98f, Align.top)

	private fun addPuzzleName(puzzle: Puzzle): GameObject = assets.addLabel(
			puzzle.name, FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.85f, Align.top)

	private fun addPuzzleText(puzzle: Puzzle): GameObject = assets.addLabel(
			puzzle.text, FontSize.F18, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.78f, Align.top)

	private fun addLevelNo(puzzle: Puzzle): GameObject = assets.addLabel(
			(level + 1).toString(), FontSize.F100, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH / 2f, Values.VIRTUAL_HEIGHT * 0.60f, Align.center)

	private fun addLeftSwipeButton(puzzle: Puzzle): GameObject = assets.addLabel(
			"{", FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH * 0.05f, Values.VIRTUAL_HEIGHT * 0.60f, Align.center)

	private fun addRightSwipeButton(puzzle: Puzzle): GameObject = assets.addLabel(
			"}", FontSize.F36, viewColor(puzzle.specs),
			Values.VIRTUAL_WIDTH * 0.95f, Values.VIRTUAL_HEIGHT * 0.60f, Align.center)

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