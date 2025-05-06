package wouldyourather

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var questionDir = ".questions"
var questionFilePath = filepath.Join(questionDir, "questions.txt")

var builtinQuestions = []string{
	"Be able to fly or Be invisible",
	"Live without music or Live without television",
	"Have super strength or Have super speed",
	"Only be able to whisper or Only be able to shout",
	"Know when you will die or Know how you will die",
	"Be rich and lonely or Be poor and surrounded by love",
	"Always be 10 minutes late or Always be 20 minutes early",
	"Eat only sweet food or Eat only salty food",
	"Have a rewind button in your life or Have a pause button",
	"Be able to talk to animals or Speak all human languages",
	"Live in space or Live under the sea",
	"Be famous or Be powerful",
	"Never have to sleep or Never have to eat",
	"Always win arguments or Always win bets",
	"Have more time or Have more money",
	"Read minds or Predict the future",
	"Be stuck on a deserted island alone or With someone you hate",
	"Have no internet or No phone",
	"Be completely bald or Covered in hair",
	"Drink only water or Only soda for the rest of your life",
	"Have to always sing instead of talk or Always dance everywhere you go",
	"Have unlimited respect or Unlimited power",
	"Lose all your money or Lose all your memories",
	"Be feared or Be loved",
	"Be a genius in a world of morons or A moron in a world of geniuses",
	"Fight one horse-sized duck or 100 duck-sized horses",
	"Have unlimited battery life on all devices or Have free WiFi everywhere",
	"Be able to teleport or Be able to time travel",
	"Always have to say what you are thinking or Never be able to speak again",
	"Have a permanent clown face or Always wear clown clothes",
	"Be able to breathe underwater or Walk through fire unscathed",
	"Be born in the past or Be born in the future",
	"Be the smartest person or The funniest person",
	"Be allergic to sunlight or Be allergic to your favorite food",
	"Never age physically or Never age mentally",
	"Be completely alone for 5 years or Constantly be surrounded by people for 5 years",
	"Only be able to lie or Only be able to tell the truth",
	"Lose the ability to read or Lose the ability to speak",
	"Have unlimited first-class tickets or Never have to pay for food at restaurants",
	"Be stuck in a horror movie or A romantic comedy",
	"Never be able to use GPS or Never be able to use a credit card",
	"Have free international flights or Never pay for a hotel again",
	"Be able to change your past or Know your future",
	"Be a superhero or Be a villain",
	"Have a pet dragon or A pet unicorn",
	"Never use social media again or Never watch TV again",
	"Only be able to whisper or Only be able to yell",
	"Have perfect memory or Be able to forget anything you want",
	"Control fire or Control water",
	"Be a famous actor or A famous musician",
	"Have a pause button for life or A skip button",
	"Own your dream home or Travel the world for free",
	"Be stuck in traffic forever or Never be able to stop moving",
	"Have no one show up to your wedding or Your funeral",
	"Be feared by all or Loved by none",
	"Sleep without dreams or Dream but never sleep",
	"Be invisible in a crowd or Be the center of attention",
	"Be rich and ugly or Poor and good-looking",
	"Never hear music again or Lose the ability to read",
	"Win the lottery or Live twice as long",
	"Be able to control animals or Control electronics with your mind",
	"Live in the city or In the countryside",
	"Have a flying car or A robot maid",
	"Be the oldest person ever or The youngest genius",
	"Have to always wear wet socks or Always wear damp clothes",
	"Have a job you love with little money or A job you hate but rich",
	"Have eyes that can record everything or Ears that can hear everything",
	"Only be able to eat one food forever or Never eat the same thing twice",
	"Sleep 1 hour a day and feel rested or Sleep 10 hours and always be tired",
	"Be super tall or Super short",
	"Be incredibly lucky or Incredibly smart",
	"Always have a full phone battery or A full tank of gas",
	"Speak every language or Play every instrument",
	"Be the best at one thing or Good at everything",
	"Live without air conditioning or Live without heating",
	"Only be able to text or Only be able to call",
	"Own a private island or A private jet",
	"Be the president or The richest person",
	"Work for someone or Work for yourself",
	"Never celebrate your birthday or Never celebrate holidays",
	"Be on a survival reality show or A dating show",
	"Have perfect skin or Perfect hair",
	"Have a third eye or A third arm",
	"Marry your best friend or Marry your celebrity crush",
	"Be stuck in your hometown forever or Never be able to return",
	"Have every movie spoiled or Never watch movies again",
	"Always be hot or Always be cold",
	"Have to always wear a suit or Always wear pajamas",
	"Be able to control your dreams or Never have nightmares",
	"Be known for something bad or Forgotten completely",
	"Win a Nobel prize or An Olympic gold medal",
	"Have the power of flight but only 1 foot above ground or Telepathy with ants",
	"Be the funniest person in the room or The most intelligent",
	"Live forever or Die tomorrow",
	"Be able to shapeshift or Turn invisible",
	"Have your dream job or Your dream partner",
	"Be stuck in the 80s or In the year 3000",
	"Be unable to lie or Always believe lies",
	"Have an endless summer or An endless winter",
	"Be able to only whisper or Only scream",
	"Lose your sense of taste or Your sense of touch",
	"Be the hero that fails or The villain that wins",
	"Have 100 good friends or 1 best friend",
	"Only wear black or Only wear white",
	"Live in a treehouse or An underground bunker",
	"Have a rewind button or A fast-forward button in life",
}

func Init() error {
	err := os.MkdirAll(questionDir, 0755)
	if err != nil {
		return err
	}
	if _, err := os.Stat(questionFilePath); os.IsNotExist(err) {
		f, err := os.Create(questionFilePath)
		if err != nil {
			return err
		}
		defer f.Close()
		for _, q := range builtinQuestions {
			f.WriteString(q + "\n")
		}
	}
	return nil
}

func LoadQuestions() ([]string, error) {
	file, err := os.Open(questionFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && strings.Contains(line, " or ") {
			questions = append(questions, line)
		}
	}
	return questions, scanner.Err()
}

func AskRandomQuestion(questions []string) error {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	parts := strings.SplitN(q, " or ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid question format")
	}
	fmt.Printf("Would you rather:\n1) %s\n2) %s\n> ", strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice != "1" && choice != "2" {
		fmt.Println("Invalid input. Please type 1 or 2.")
		return nil
	}

	first := rand.Intn(101)
	second := 100 - first
	if choice == "2" {
		first, second = second, first
	}

	fmt.Printf("You chose option %s.\nStats: 1) %d%%  2) %d%%\n\n", choice, first, second)
	return nil
}
