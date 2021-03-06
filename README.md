![Kalbi Logo](https://raw.githubusercontent.com/hyperioxx/Kalbi/master/assets/images/logo_transparent_background.png "Kalbi Logo")

# Kalbi - Golang Session Initiated Protocol Framework  


[![Go Report Card](https://goreportcard.com/badge/github.com/hyperioxx/Kalbi)](https://goreportcard.com/report/github.com/hyperioxx/Kalbi) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/KalbiProject/Kalbi/LICENCE)  [![GitHub contributors](https://img.shields.io/github/contributors/Naereen/StrapDown.js.svg)](https://github.com/KalbiProject/Kalbi/graphs/contributors/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Hyperioxx/Kalbi) [![Discord](https://discordapp.com/api/guilds/762641136425762816/widget.png)](https://discord.gg/6NCKgrz) 
![GitHub Repo stars](https://img.shields.io/github/stars/Hyperioxx/Kalbi?style=social)  ![GitHub forks](https://img.shields.io/github/forks/hyperioxx/Kalbi?style=social)
![Subreddit subscribers](https://img.shields.io/reddit/subreddit-subscribers/Kalbi?style=social)

## Description

Golang SIP/VoIP SDK's used to build large platforms for VoIP and realtime communications

## Join The Discord

Discord - https://discord.gg/6NCKgrz

# High Level Architecture View

Withing your application create a struct following the EventListener interface and register it to the SIPStack either with a go channel or a callback function for SIP Events

![Basic View](https://raw.githubusercontent.com/hyperioxx/Kalbi/master/doc/BasicView.png "Basic View")

