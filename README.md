mh4gquest
=========

亂數挑選 MH4G 任務，包含武器及特殊限制

Usage
=====

	mh4gquest -file=quests.json

Docker
======

Build Image
-----------

	docker build -t <image_tag> -f ./Dockerfile .

Run Container
-------------

	docker run -d -p 8080:8080 <image_tag>

TODO
====

* 任務等級挑選
* 武器限制
* 處理網頁要求
