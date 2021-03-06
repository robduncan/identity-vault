/*
 * Copyright (C) 2016-2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
'use strict'
var Ajax = require('./Ajax');

var Model = {
	url: 'models',

	list: function () {
			return Ajax.get(this.url);
	},

	get: function(modelId) {
		return Ajax.get(this.url + '/' + modelId);
	},

	update:  function(model) {
		return Ajax.put(this.url + '/' + model.id, model);
	},

	delete:  function(model) {
		return Ajax.delete(this.url + '/' + model.id, {});
	},

	create:  function(model) {
		return Ajax.post(this.url, model);
	}

}

module.exports = Model;
