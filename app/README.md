# ytlive

An app to find what cool channels are live right now

## YoutubeAPI

Package youtubeAPI defines all the methods with which this app interacts with the youtube DATA API.

## Server

Package server defines all the handler functions that can be called by the client.

## Store

Package store provides an interface with which we can store data, be it to a file or DB.

## Checker

Package checker defines the checker agent that keeps track of the current status of known live streams.
e.g. current view count and live status

## Getter

Package getter defines methods with which the app can scrape Youtube for data that is unavailable, such as Category
