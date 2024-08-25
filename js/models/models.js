// Generated types for models
/**
 * @typedef {Object} Location
 * @property {string} venue
 * @property {string} address
 * @property {string} city
 * @property {string} country
 * @property {Coordinates} coordinates [optional]
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
 * @typedef {Object} Sessions
 * @property {string} title
 * @property {string} description
 * @property {Date} start_time
 * @property {Date} end_time
 * @property {string} room [optional]
 * @property {Speaker} speaker
 * @property {string} id
 */
class Sessions {
  constructor(data) {
    this.start_time = data.start_time;
    this.end_time = data.end_time;
    this.room = data.room;
    this.speaker = data.speaker;
    this.id = data.id;
    this.title = data.title;
    this.description = data.description;
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
 * @typedef {Object} BlogPost
 * @property {SEOMetadata} seometadata [optional]
 * @property {Comment[]} comment [optional]
 * @property {string} id
 * @property {string} author_id
 * @property {string} status
 * @property {Date} updated_at
 * @property {string[]} categories [optional]
 * @property {string[]} tags [optional]
 * @property {string} title
 * @property {string} content
 * @property {Date} created_at
 */
class BlogPost {
  constructor(data) {
    this.id = data.id;
    this.author_id = data.author_id;
    this.status = data.status;
    this.seometadata = data.seometadata;
    this.comment = data.comment;
    this.title = data.title;
    this.content = data.content;
    this.created_at = data.created_at;
    this.updated_at = data.updated_at;
    this.categories = data.categories;
    this.tags = data.tags;
  }
}

/**
 * @typedef {Object} Sponsors
 * @property {string} id
 * @property {string} name
 * @property {string} logo_url
 * @property {string} sponsorship_level
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
 * @typedef {Object} Vip
 * @property {number} available_quantity
 * @property {string} perks
 * @property {string} id
 * @property {string} name
 * @property {number} price
 */
class Vip {
  constructor(data) {
    this.price = data.price;
    this.available_quantity = data.available_quantity;
    this.perks = data.perks;
    this.id = data.id;
    this.name = data.name;
  }
}

/**
 * @typedef {Object} Replies
 * @property {string} user_id
 * @property {string} content
 * @property {Date} created_at
 * @property {string} id
 */
class Replies {
  constructor(data) {
    this.id = data.id;
    this.user_id = data.user_id;
    this.content = data.content;
    this.created_at = data.created_at;
  }
}

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
 * @typedef {Object} Comment
 * @property {string} id
 * @property {string} user_id
 * @property {string} content
 * @property {Date} created_at
 * @property {Replies[]} replies [optional]
 */
class Comment {
  constructor(data) {
    this.id = data.id;
    this.user_id = data.user_id;
    this.content = data.content;
    this.created_at = data.created_at;
    this.replies = data.replies;
  }
}

/**
 * @typedef {Object} Event
 * @property {Date} end_time
 * @property {Location} location
 * @property {Organizer} organizer
 * @property {Ticket_types} ticket_types
 * @property {Date} start_time
 * @property {string} name
 * @property {string} description
 * @property {Sponsors[]} sponsors [optional]
 * @property {Sessions[]} sessions [optional]
 * @property {string} id
 */
class Event {
  constructor(data) {
    this.description = data.description;
    this.sponsors = data.sponsors;
    this.sessions = data.sessions;
    this.id = data.id;
    this.name = data.name;
    this.location = data.location;
    this.organizer = data.organizer;
    this.ticket_types = data.ticket_types;
    this.start_time = data.start_time;
    this.end_time = data.end_time;
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
 * @property {string} id
 * @property {string} name
 * @property {number} price
 * @property {number} available_quantity
 */
class Standard {
  constructor(data) {
    this.id = data.id;
    this.name = data.name;
    this.price = data.price;
    this.available_quantity = data.available_quantity;
  }
}

/**
 * @typedef {Object} SEOMetadata
 * @property {string} meta_description
 * @property {string} canonical_url
 * @property {string} meta_title
 */
class SEOMetadata {
  constructor(data) {
    this.meta_title = data.meta_title;
    this.meta_description = data.meta_description;
    this.canonical_url = data.canonical_url;
  }
}


// Export all types
export {
    Location,
    Sessions,
    Speaker,
    BlogPost,
    Sponsors,
    Vip,
    Replies,
    Coordinates,
    Ticket_types,
    Comment,
    Event,
    Organizer,
    Standard,
    SEOMetadata,
}

// Types generated by type-utils

