// Generated types for models
/**
 * @typedef {Object} Coordinates
 * @property {number} latitude
 * @property {number} longitude
 */
class Coordinates {
  constructor(data) {
    this.latitude = data.latitude;
    this.longitude = data.longitude;
  }
}

/**
 * @typedef {Object} Event
 * @property {string} name
 * @property {Organizer} organizer
 * @property {string} id
 * @property {Date} start_time
 * @property {Date} end_time
 * @property {Location} location
 * @property {Sponsors[]} sponsors [optional]
 * @property {Sessions[]} sessions [optional]
 * @property {Ticket_types} ticket_types
 * @property {string} description
 */
class Event {
  constructor(data) {
    this.description = data.description;
    this.start_time = data.start_time;
    this.end_time = data.end_time;
    this.location = data.location;
    this.sponsors = data.sponsors;
    this.sessions = data.sessions;
    this.ticket_types = data.ticket_types;
    this.id = data.id;
    this.name = data.name;
    this.organizer = data.organizer;
  }
}

/**
 * @typedef {Object} Organizer
 * @property {string} id
 * @property {string} name
 * @property {string} email
 */
class Organizer {
  constructor(data) {
    this.id = data.id;
    this.name = data.name;
    this.email = data.email;
  }
}

/**
 * @typedef {Object} Standard
 * @property {number} available_quantity
 * @property {string} id
 * @property {string} name
 * @property {number} price
 */
class Standard {
  constructor(data) {
    this.price = data.price;
    this.available_quantity = data.available_quantity;
    this.id = data.id;
    this.name = data.name;
  }
}

/**
 * @typedef {Object} BlogPost
 * @property {string} id
 * @property {string} author_id
 * @property {string[]} tags [optional]
 * @property {string} title
 * @property {string} content
 * @property {Date} created_at
 * @property {Date} updated_at
 * @property {string} status
 * @property {string[]} categories [optional]
 * @property {SEOMetadata} seometadata [optional]
 * @property {Comment[]} comment [optional]
 */
class BlogPost {
  constructor(data) {
    this.updated_at = data.updated_at;
    this.status = data.status;
    this.categories = data.categories;
    this.seometadata = data.seometadata;
    this.comment = data.comment;
    this.title = data.title;
    this.content = data.content;
    this.created_at = data.created_at;
    this.id = data.id;
    this.author_id = data.author_id;
    this.tags = data.tags;
  }
}

/**
 * @typedef {Object} Comment
 * @property {string} user_id
 * @property {string} content
 * @property {Date} created_at
 * @property {Replies[]} replies [optional]
 * @property {string} id
 */
class Comment {
  constructor(data) {
    this.created_at = data.created_at;
    this.replies = data.replies;
    this.id = data.id;
    this.user_id = data.user_id;
    this.content = data.content;
  }
}

/**
 * @typedef {Object} Replies
 * @property {Date} created_at
 * @property {string} id
 * @property {string} user_id
 * @property {string} content
 */
class Replies {
  constructor(data) {
    this.user_id = data.user_id;
    this.content = data.content;
    this.created_at = data.created_at;
    this.id = data.id;
  }
}

/**
 * @typedef {Object} Sessions
 * @property {string} room [optional]
 * @property {Speaker} speaker
 * @property {string} id
 * @property {string} title
 * @property {string} description
 * @property {Date} start_time
 * @property {Date} end_time
 */
class Sessions {
  constructor(data) {
    this.title = data.title;
    this.description = data.description;
    this.start_time = data.start_time;
    this.end_time = data.end_time;
    this.room = data.room;
    this.speaker = data.speaker;
    this.id = data.id;
  }
}

/**
 * @typedef {Object} SEOMetadata
 * @property {string} canonical_url
 * @property {string} meta_title
 * @property {string} meta_description
 */
class SEOMetadata {
  constructor(data) {
    this.meta_description = data.meta_description;
    this.canonical_url = data.canonical_url;
    this.meta_title = data.meta_title;
  }
}

/**
 * @typedef {Object} Location
 * @property {string} city
 * @property {string} country
 * @property {Coordinates} coordinates [optional]
 * @property {string} venue
 * @property {string} address
 */
class Location {
  constructor(data) {
    this.address = data.address;
    this.city = data.city;
    this.country = data.country;
    this.coordinates = data.coordinates;
    this.venue = data.venue;
  }
}

/**
 * @typedef {Object} Sponsors
 * @property {string} logo_url
 * @property {string} sponsorship_level
 * @property {string} id
 * @property {string} name
 */
class Sponsors {
  constructor(data) {
    this.id = data.id;
    this.name = data.name;
    this.logo_url = data.logo_url;
    this.sponsorship_level = data.sponsorship_level;
  }
}

/**
 * @typedef {Object} Speaker
 * @property {string} id
 * @property {string} name
 * @property {string} bio
 */
class Speaker {
  constructor(data) {
    this.id = data.id;
    this.name = data.name;
    this.bio = data.bio;
  }
}

/**
 * @typedef {Object} Ticket_types
 * @property {Standard[]} standard
 * @property {Vip[]} vip [optional]
 */
class Ticket_types {
  constructor(data) {
    this.standard = data.standard;
    this.vip = data.vip;
  }
}

/**
 * @typedef {Object} Vip
 * @property {string} id
 * @property {string} name
 * @property {number} price
 * @property {number} available_quantity
 * @property {string} perks
 */
class Vip {
  constructor(data) {
    this.id = data.id;
    this.name = data.name;
    this.price = data.price;
    this.available_quantity = data.available_quantity;
    this.perks = data.perks;
  }
}


// CommonJS module exports
module.exports = {
    Coordinates,
    Event,
    Organizer,
    Standard,
    BlogPost,
    Comment,
    Replies,
    Sessions,
    SEOMetadata,
    Location,
    Sponsors,
    Speaker,
    Ticket_types,
    Vip
};




// Types generated by type-utils

