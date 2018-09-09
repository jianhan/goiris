package zomato

type CategoryResponse struct {
	Categories []struct {
		Categories struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"categories"`
	} `json:"categories"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CitiesResponse struct {
	LocationSuggestions []City `json:"location_suggestions"`
	Status              string `json:"status"`
	HasMore             int    `json:"has_more"`
	HasTotal            int    `json:"has_total"`
}

type City struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	CountryID            int    `json:"country_id"`
	CountryName          string `json:"country_name"`
	CountryFlagURL       string `json:"country_flag_url"`
	ShouldExperimentWith int    `json:"should_experiment_with"`
	DiscoveryEnabled     int    `json:"discovery_enabled"`
	HasNewAdFormat       int    `json:"has_new_ad_format"`
	IsState              int    `json:"is_state"`
	StateID              int    `json:"state_id"`
	StateName            string `json:"state_name"`
	StateCode            string `json:"state_code"`
}

type Collection struct {
	CollectionID string `json:"collection_id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`
	ResCount     string `json:"res_count"`
	ShareURL     string `json:"share_url"`
}

type Cuisine struct {
	CuisineID   string `json:"cuisine_id"`
	CuisineName string `json:"cuisine_name"`
}

type Establishment struct {
	EstablishmentID   string `json:"establishment_id"`
	EstablishmentName string `json:"establishment_name"`
}

type GeoCode struct {
	Locality struct {
		EntityType  string `json:"entity_type"`
		EntityID    string `json:"entity_id"`
		Title       string `json:"title"`
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
		CityID      string `json:"city_id"`
		CityName    string `json:"city_name"`
		CountryID   string `json:"country_id"`
		CountryName string `json:"country_name"`
	} `json:"locality"`
	Popularity struct {
		Popularity     string   `json:"popularity"`
		NightlifeIndex string   `json:"nightlife_index"`
		TopCuisines    []string `json:"top_cuisines"`
	} `json:"popularity"`
	Link              string `json:"link"`
	NearbyRestaurants []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		Location struct {
			Address   string `json:"address"`
			Locality  string `json:"locality"`
			City      string `json:"city"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
			Zipcode   string `json:"zipcode"`
			CountryID string `json:"country_id"`
		} `json:"location"`
		AverageCostForTwo string `json:"average_cost_for_two"`
		PriceRange        string `json:"price_range"`
		Currency          string `json:"currency"`
		Thumb             string `json:"thumb"`
		FeaturedImage     string `json:"featured_image"`
		PhotosURL         string `json:"photos_url"`
		MenuURL           string `json:"menu_url"`
		EventsURL         string `json:"events_url"`
		UserRating        struct {
			AggregateRating string `json:"aggregate_rating"`
			RatingText      string `json:"rating_text"`
			RatingColor     string `json:"rating_color"`
			Votes           string `json:"votes"`
		} `json:"user_rating"`
		HasOnlineDelivery string `json:"has_online_delivery"`
		IsDeliveringNow   string `json:"is_delivering_now"`
		HasTableBooking   string `json:"has_table_booking"`
		Deeplink          string `json:"deeplink"`
		Cuisines          string `json:"cuisines"`
		AllReviewsCount   string `json:"all_reviews_count"`
		PhotoCount        string `json:"photo_count"`
		PhoneNumbers      string `json:"phone_numbers"`
		Photos            []struct {
			ID       string `json:"id"`
			URL      string `json:"url"`
			ThumbURL string `json:"thumb_url"`
			User     struct {
				Name            string `json:"name"`
				ZomatoHandle    string `json:"zomato_handle"`
				FoodieLevel     string `json:"foodie_level"`
				FoodieLevelNum  string `json:"foodie_level_num"`
				FoodieColor     string `json:"foodie_color"`
				ProfileURL      string `json:"profile_url"`
				ProfileDeeplink string `json:"profile_deeplink"`
				ProfileImage    string `json:"profile_image"`
			} `json:"user"`
			ResID         string `json:"res_id"`
			Caption       string `json:"caption"`
			Timestamp     string `json:"timestamp"`
			FriendlyTime  string `json:"friendly_time"`
			Width         string `json:"width"`
			Height        string `json:"height"`
			CommentsCount string `json:"comments_count"`
			LikesCount    string `json:"likes_count"`
		} `json:"photos"`
		AllReviews []struct {
			Rating             string `json:"rating"`
			ReviewText         string `json:"review_text"`
			ID                 string `json:"id"`
			RatingColor        string `json:"rating_color"`
			ReviewTimeFriendly string `json:"review_time_friendly"`
			RatingText         string `json:"rating_text"`
			Timestamp          string `json:"timestamp"`
			Likes              string `json:"likes"`
			User               struct {
				Name            string `json:"name"`
				ZomatoHandle    string `json:"zomato_handle"`
				FoodieLevel     string `json:"foodie_level"`
				FoodieLevelNum  string `json:"foodie_level_num"`
				FoodieColor     string `json:"foodie_color"`
				ProfileURL      string `json:"profile_url"`
				ProfileDeeplink string `json:"profile_deeplink"`
				ProfileImage    string `json:"profile_image"`
			} `json:"user"`
			CommentsCount string `json:"comments_count"`
		} `json:"all_reviews"`
	} `json:"nearby_restaurants"`
}
