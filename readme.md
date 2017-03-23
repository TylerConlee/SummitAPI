# SummitAPI

## Introduction

SummitAPI is an automated tool used to pull data from Google
Analytics, Google Adwords, Bing Ads, SEMRush, and more. With SummitAPI,
the aggregated data from these sources is then able to be manipulated
and viewed utilizing the Summit Dashboard.

Currently, the API pulls information from:

- Google Analytics

Currently, the API has plans to also include information from:

- Google AdWords
- Bing Ads
- SEMRush
- New Relic


Data is stored locally within Cassandra, with individual data
points for a daily basis. Initial data points, used as a proof of
concept, will be basic, but allow for data manipulation over multiple
time windows.


## Installation Instructions

To install and run the SummitAPI locally, clone the repository and build
using Go 1.7 and above. As the project matures, this will eventually be
a Dockerfile based installation, where the only thing needed would be to
pull and run the associated Docker image. It'll get there. Someday.

## Additional Resources