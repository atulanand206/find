package com.creations.games.think.scenes.main.objects

import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.data.Puzzle
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.BackgroundRounded
import com.creations.games.think.scenes.DarkThemeColor
import com.creations.games.think.utils.assets

class PuzzleFrame(private val di: DI): GameObject(), DarkThemeColor {

	private lateinit var backgroundRounded: GameObject
	private lateinit var puzzleName: GameObject
	private lateinit var puzzleText: GameObject
	private lateinit var levelSwipe: GameObject
	private var level = 0

	init {
		addFrameBackground()
		addPuzzle(xTheme.puzzles[level])
		addLevel()
	}

	private fun addFrameBackground() {
		backgroundRounded = BackgroundRounded(di)
		backgroundRounded.setPosition(x, y)
		backgroundRounded.color = viewColorInvert(xTheme.specs)
		addActor(backgroundRounded)
	}

	private fun addPuzzle(puzzle: Puzzle) {
		puzzleName = addPuzzleName(puzzle)
		puzzleText = addPuzzleText(puzzle)
		addActor(puzzleName)
		addActor(puzzleText)
	}

	private fun addLevel() {
		levelSwipe = LevelSwiper(di, level)
		levelSwipe.setPosition(x, y + Values.VIRTUAL_HEIGHT*0.56f)
		addActor(levelSwipe)
	}

	private fun addPuzzleName(puzzle: Puzzle): GameObject = di.assets.addWrappedLabel(
			puzzle.name, FontSize.F36, viewColor(puzzle.specs),
			x + Values.VIRTUAL_WIDTH*0.5f, y + Values.VIRTUAL_HEIGHT*0.85f, Align.center)

	private fun addPuzzleText(puzzle: Puzzle): GameObject = di.assets.addWrappedLabel(
			puzzle.text, FontSize.F18, viewColor(puzzle.specs),
			x + Values.VIRTUAL_WIDTH*0.5f, y + Values.VIRTUAL_HEIGHT*0.72f, Align.center)

}