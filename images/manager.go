package images

import (
        "fmt"
        "friday-bot/logger"
        "math/rand"
        "os"
        "path/filepath"
        "strings"
        "time"
)

// Manager handles image operations
type Manager struct {
        imagesDir string
        random    *rand.Rand
}

// New creates a new image manager
func New(imagesDir string) *Manager {
        return &Manager{
                imagesDir: imagesDir,
                random:    rand.New(rand.NewSource(time.Now().UnixNano())),
        }
}

// GetRandomImage returns a random image path from the images directory
func (m *Manager) GetRandomImage() (string, error) {
        images, err := m.listImages()
        if err != nil {
                return "", fmt.Errorf("failed to list images: %w", err)
        }

        if len(images) == 0 {
                return "", fmt.Errorf("no images found in directory: %s", m.imagesDir)
        }

        // Select random image
        selectedImage := images[m.random.Intn(len(images))]
        logger.Info("Selected image: %s", selectedImage)
        
        return selectedImage, nil
}

// listImages returns all image files in the images directory
func (m *Manager) listImages() ([]string, error) {
        // Create directory if it doesn't exist
        if err := os.MkdirAll(m.imagesDir, 0755); err != nil {
                return nil, fmt.Errorf("failed to create images directory: %w", err)
        }

        var images []string
        supportedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg"}

        err := filepath.Walk(m.imagesDir, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        return err
                }

                if info.IsDir() {
                        return nil
                }

                // Check if file has supported extension
                ext := strings.ToLower(filepath.Ext(path))
                for _, supportedExt := range supportedExtensions {
                        if ext == supportedExt {
                                images = append(images, path)
                                break
                        }
                }

                return nil
        })

        if err != nil {
                return nil, fmt.Errorf("failed to walk images directory: %w", err)
        }

        logger.Info("Found %d images in directory %s", len(images), m.imagesDir)
        return images, nil
}

// GetImageCount returns the number of available images
func (m *Manager) GetImageCount() (int, error) {
        images, err := m.listImages()
        if err != nil {
                return 0, err
        }
        return len(images), nil
}

// ValidateImagesDirectory checks if the images directory exists and contains images
func (m *Manager) ValidateImagesDirectory() error {
        if _, err := os.Stat(m.imagesDir); os.IsNotExist(err) {
                return fmt.Errorf("images directory does not exist: %s", m.imagesDir)
        }

        count, err := m.GetImageCount()
        if err != nil {
                return fmt.Errorf("failed to count images: %w", err)
        }

        if count == 0 {
                return fmt.Errorf("no images found in directory: %s", m.imagesDir)
        }

        return nil
}
