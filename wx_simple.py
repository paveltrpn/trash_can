#!/usr/bin/python

import wx

def main():
	app = wx.App()

	frame = wx.Frame(None, -1, 'simple.py')
	frame.Show()

	app.MainLoop()
	
if __name__ == "__main__":
	main()