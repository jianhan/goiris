package zomato

type DailyMenu struct {
	DailyMenuID string `json:"daily_menu_id"`
	Name        string `json:"name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Dishes      []struct {
		DishID string `json:"dish_id"`
		Name   string `json:"name"`
		Price  string `json:"price"`
	} `json:"dishes"`
}

type Restaurant struct {
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
}

type Review struct {
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
}

type Search struct {
	ResultsFound string `json:"results_found"`
	ResultsStart string `json:"results_start"`
	ResultsShown string `json:"results_shown"`
	Restaurants  []struct {
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
	} `json:"restaurants"`
}
