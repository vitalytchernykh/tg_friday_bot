package scheduler

import (
	"friday-bot/bot"
	"friday-bot/config"
	"friday-bot/logger"
	"time"
)

// Scheduler handles the Friday posting schedule
type Scheduler struct {
	bot         *bot.Bot
	config      *config.Config
	ticker      *time.Ticker
	stopChannel chan struct{}
	lastPosted  time.Time
}

// New creates a new Scheduler instance
func New(bot *bot.Bot, cfg *config.Config) *Scheduler {
	return &Scheduler{
		bot:         bot,
		config:      cfg,
		stopChannel: make(chan struct{}),
	}
}

// Start begins the scheduling loop
func (s *Scheduler) Start() {
	// Create ticker for the check interval
	s.ticker = time.NewTicker(time.Duration(s.config.CheckInterval) * time.Minute)
	defer s.ticker.Stop()

	logger.Info("Scheduler started. Checking every %d minutes for Friday posts", s.config.CheckInterval)

	// Check immediately on startup
	s.checkAndPost()

	for {
		select {
		case <-s.ticker.C:
			s.checkAndPost()
		case <-s.stopChannel:
			logger.Info("Scheduler stopped")
			return
		}
	}
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	close(s.stopChannel)
	if s.ticker != nil {
		s.ticker.Stop()
	}
}

// checkAndPost checks if it's time to post and posts if needed
func (s *Scheduler) checkAndPost() {
	now := time.Now()
	
	// Check if today is Friday
	if now.Weekday() != time.Friday {
		logger.Debug("Today is %s, not Friday. Skipping post check.", now.Weekday().String())
		return
	}

	// Check if we're at the right time
	if !s.isPostTime(now) {
		logger.Debug("Not yet post time. Current: %02d:%02d, Target: %02d:%02d", 
			now.Hour(), now.Minute(), s.config.PostHour, s.config.PostMinute)
		return
	}

	// Check if we already posted today
	if s.alreadyPostedToday(now) {
		logger.Debug("Already posted today. Last posted: %s", s.lastPosted.Format("2006-01-02 15:04:05"))
		return
	}

	// Time to post!
	logger.Info("It's Friday and time to post! Posting image...")
	
	if err := s.bot.PostFridayImage(s.config.ChatID); err != nil {
		logger.Error("Failed to post Friday image: %v", err)
		return
	}

	// Update last posted time
	s.lastPosted = now
	logger.Info("Successfully posted Friday image at %s", now.Format("2006-01-02 15:04:05"))
}

// isPostTime checks if current time matches the configured post time
func (s *Scheduler) isPostTime(now time.Time) bool {
	currentHour := now.Hour()
	currentMinute := now.Minute()
	
	// Check if we're within the post time window (allowing some flexibility)
	targetHour := s.config.PostHour
	targetMinute := s.config.PostMinute
	
	// Allow posting within a window around the target time
	if currentHour == targetHour {
		// If we're in the same hour, check minutes
		return currentMinute >= targetMinute && currentMinute < targetMinute+s.config.CheckInterval
	}
	
	// If check interval spans across hours, handle that case
	if s.config.CheckInterval > 60 {
		windowEnd := time.Date(now.Year(), now.Month(), now.Day(), targetHour, targetMinute, 0, 0, now.Location()).
			Add(time.Duration(s.config.CheckInterval) * time.Minute)
		
		targetTime := time.Date(now.Year(), now.Month(), now.Day(), targetHour, targetMinute, 0, 0, now.Location())
		
		return now.After(targetTime) && now.Before(windowEnd)
	}
	
	return false
}

// alreadyPostedToday checks if we already posted today
func (s *Scheduler) alreadyPostedToday(now time.Time) bool {
	if s.lastPosted.IsZero() {
		return false
	}
	
	// Check if last post was today
	lastYear, lastMonth, lastDay := s.lastPosted.Date()
	nowYear, nowMonth, nowDay := now.Date()
	
	return lastYear == nowYear && lastMonth == nowMonth && lastDay == nowDay
}

// GetNextFridayTime returns the next Friday posting time
func (s *Scheduler) GetNextFridayTime() time.Time {
	now := time.Now()
	
	// Calculate days until next Friday
	daysUntilFriday := (int(time.Friday) - int(now.Weekday()) + 7) % 7
	if daysUntilFriday == 0 {
		// Today is Friday, check if we should post today or next Friday
		postTime := time.Date(now.Year(), now.Month(), now.Day(), s.config.PostHour, s.config.PostMinute, 0, 0, now.Location())
		if now.After(postTime) {
			// Already passed today's post time, next Friday
			daysUntilFriday = 7
		}
	}
	
	nextFriday := now.AddDate(0, 0, daysUntilFriday)
	return time.Date(nextFriday.Year(), nextFriday.Month(), nextFriday.Day(), 
		s.config.PostHour, s.config.PostMinute, 0, 0, now.Location())
}
