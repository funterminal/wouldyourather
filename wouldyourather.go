package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	questionDir      = ".questions"
	questionFilePath = filepath.Join(questionDir, "questions.txt")
)

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
	"Have the ability to change the world but at a personal cost or Leave the world as it is but never make an impact",
	"Always know when someone is lying or Never be lied to",
	"Find your soulmate but never see them again or Never find your soulmate but live a happy life",
	"Be able to survive any disaster or Stop any disaster from happening",
	"Live in a world without disease or Without crime",
	"Experience the perfect day or the perfect week",
	"Always be happy but never successful or Always be successful but never happy",
	"Live with a secret superpower that no one knows or Have an obvious power everyone knows about",
	"Be able to pause time but only for 10 seconds or Be able to rewind time but only by 10 minutes",
	"Be a time traveler but never be able to change history or Never time travel but know all the secrets of time",
	"Always have the perfect answer to any question or Always have the perfect response in any conversation",
	"Have the ability to talk to animals or the ability to control the weather",
	"Have the ability to change your appearance or the ability to change your voice",
	"Never feel tired again or never feel hungry again",
	"Be able to instantly master any skill or Have unlimited access to any book or knowledge",
	"Always be completely honest or Always be completely diplomatic",
	"Be able to breathe underwater or be able to fly for short distances",
	"Always feel calm in stressful situations or Always have a clear and sharp mind when making decisions",
	"Have the ability to erase memories or Have the ability to read minds",
	"Know when you are about to die or Never know the day of your death",
	"Always be lucky but have bad health or Never be lucky but be in perfect health",
	"Have your future perfectly planned out or Live completely spontaneously without knowing what’s next",
	"Have a life that is always interesting or A life that is always peaceful",
	"Be able to stay young forever but never fall in love or Be able to age normally but experience true love",
	"Have everything you ever wanted but lose all your relationships or Have an average life but have all the love and friendship you need",
	"Live in a world without social media or Live in a world without phones",
	"Have the ability to erase any painful memory or Have the ability to relive any happy memory",
	"Only be able to talk in song lyrics or Only be able to talk in movie quotes",
	"Always live in the moment or Always have the perfect long-term plan",
	"Have a superpower but only when you’re alone or Never have a superpower but always be surrounded by superheroes",
	"Be able to speak all languages or Be able to communicate with animals",
	"Have the power to time travel or the power to read minds",
	"Live without technology or Live without human contact",
	"Find true love or Become incredibly successful",
	"Have unlimited wealth or Unlimited knowledge",
	"Never age or Never need sleep",
	"Have the power to fly but always be cold or Be invisible but always feel warm",
	"Never feel pain again or Never feel sadness again",
	"Win a million dollars but never see your family again or Lose everything but always have your family by your side",
	"Have a pet dinosaur or a pet dragon",
	"Live in the future or live in the past",
	"Be the richest person on earth but never have true friends or Have deep friendships but be poor",
	"Live forever but alone or Live a short life full of love and adventure",
	"Be stuck in a video game or A movie for the rest of your life",
	"Have the ability to teleport but only to places you've already been or Be able to teleport anywhere but only once",
	"Live in a house made of candy or A house made of ice",
	"Have the ability to make anyone tell the truth or Make anyone fall in love with you",
	"Be famous for something bad or Be unknown but loved by everyone",
	"Be able to read any book instantly or Learn a new skill instantly",
	"Have the ability to never get tired or Always have the perfect amount of rest",
	"Have the power to heal others but not yourself or Heal yourself but never others",
	"Live in a world with no music or A world with no art",
	"Never be able to use the internet again or Never be able to leave your country",
	"Always have to tell the truth or Always lie and never get caught",
	"Be able to talk to plants or Be able to talk to machines",
	"Have the ability to make any food appear instantly or The ability to never gain weight no matter what you eat",
	"Never use social media again or Never watch TV again",
	"Be able to erase memories or Be able to relive one memory forever",
	"Always be feared but never loved or Always be loved but never respected",
	"Be able to control time but only for 10 seconds or Never control time but always know the future",
	"Have unlimited money but never find happiness or Live a modest life but always be happy",
	"Only be able to watch one TV show forever or Only listen to one song forever",
	"Always be 10 minutes late or Always be 10 minutes early",
	"Have to sing every time you speak or Dance every time you move",
	"Never be able to use a phone again or Never be able to use a computer again",
	"Be able to read people's thoughts or Be able to see their future",
	"Always be misunderstood but never explain yourself or Always be understood but never get your point across",
	"Be feared by everyone or Be loved by everyone but only be a background character in life",
	"Live without ever experiencing failure or Live without ever experiencing success",
	"Always be invisible but no one knows or Always be visible but never be noticed",
	"Have the ability to always know the truth but never be able to say it or Always be able to lie but never tell the truth",
	"Live in a world where everyone is honest but lonely or A world full of lies but filled with friends",
	"Be able to time travel to the past but never return or Time travel to the future and never be able to go back",
	"Have to spend the rest of your life on an island with no technology or In a city full of technology but no nature",
	"Be able to stop time but only for an hour a day or Have perfect control of your environment but no time control",
	"Live your life in a perfect bubble or Always face danger but be free",
	"Have a memory that never fades or Forget everything but always be happy",
	"Only be able to wear one outfit for the rest of your life or Only be able to change your outfit once a year",
	"Be able to speak to animals or Be able to understand plants",
	"Be forced to sing every time you speak or Dance every time you walk",
	"Be an astronaut but never return to earth or Live on earth but never travel again",
	"Live without internet or Live without your phone",
	"Only be able to communicate through text or Always have to speak but never use words",
	"Have superhuman intelligence or Superhuman strength",
	"Always be able to find the truth or Always be able to find a solution",
	"Be able to stop any crime but only at the cost of your own safety or Never be able to stop a crime but save yourself each time",
	"Be a genius in a world of idiots or A fool in a world of geniuses",
	"Always know when someone is lying to you or Always be able to tell when someone is telling the truth",
	"Live in a world where everyone is the same or A world full of completely unique individuals",
	"Have super speed but constantly be tired or Super strength but always feel weak",
	"Be able to change any part of your body at will or Change your environment to your liking",
	"Be able to live forever but always be alone or Live a normal life but always surrounded by love",
	"Have all the knowledge in the world but no ability to share it or Never learn anything new but be able to help others",
	"Be able to control the weather but only in one small area or Control the weather anywhere but with limited power",
	"Be the richest person alive but no one respects you or The poorest but everyone loves you",
	"Only be able to walk on your hands or Only walk backwards",
	"Never be able to sleep or Never be able to eat",
	"Have your dream job but never get a promotion or Have a regular job but constantly move up the ladder",
	"Live without love or Live without money",
	"Have the power to read minds but never know your own thoughts or Never read minds but always know your own mind",
	"Live a simple life in the countryside or Live a complicated life in a busy city",
	"Have a house that can never be messy or Always have a messy house but never have to clean it",
	"Always be the smartest person in the room but never fit in or Always be the least smart but loved by everyone",
	"Be able to talk to everyone but only in riddles or Never talk but always communicate perfectly",
	"Never be able to leave your hometown or Travel the world but never stay in one place",
	"Live your life as a hero but always sacrifice or Live as a villain but always be loved",
	"Live in a utopia but never be happy or Live in a dystopia but always feel fulfilled",
	"Be able to instantly learn any language or Learn one language in-depth and master it",
	"Never have to wait in line or Never have to pay for anything ever again",
	"Always be able to travel for free but never have a permanent home or Have a permanent home but never leave your city",
	"Have the ability to read every book instantly or Only be able to read one book for the rest of your life",
	"Always have perfect skin or Always have perfect hair",
	"Be able to live in any fictional universe but never return home or Never visit a fictional world but always have a perfect life",
	"Be able to change one thing about your past or Know everything that happens in your future",
	"Always be in control of your emotions or Never feel regret",
	"Be able to fly but only a few inches above the ground or Have super strength but only for an hour a day",
	"Always feel safe but never have an adventure or Always have an adventure but never feel safe",
	"Never be able to speak or Never be able to hear",
	"Always feel unimportant or Always feel like everyone is watching you",
	"Live a life of extreme adventure or a life of quiet simplicity",
	"Be able to read every person's thoughts or Never know what anyone is thinking",
	"Be feared by everyone or Loved by everyone but never respected",
	"Be the hero of a story that never ends or The villain of a story that ends happily for others",
	"Always be underestimated or Always be overestimated",
	"Have a perfect memory but never forget anything or Forget all your painful memories but lose the lessons learned",
	"Live forever in peace but never achieve anything or Live a short life of adventure and accomplishment",
	"Be incredibly lucky but never make a lasting impact or Be extremely unlucky but change the world",
	"Be the greatest lover or the greatest friend",
	"Only be able to text or Only be able to speak in person",
	"Never experience fear or Never experience regret",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wouldyourather init | ask")
		return
	}
	switch os.Args[1] {
	case "init":
		err := Init()
		if err != nil {
			fmt.Println("Error initializing:", err)
		} else {
			fmt.Println("Question database initialized.")
		}
	case "ask":
		questions, err := LoadQuestions()
		if err != nil {
			fmt.Println("Error loading questions:", err)
			return
		}
		for {
			err = AskRandomQuestion(questions)
			if err != nil {
				fmt.Println("Error asking question:", err)
				return
			}
		}
	default:
		fmt.Println("Unknown command. Use: init | ask")
	}
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
			f.WriteString(q + "\n") // Fixed newline handling
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
	choice, _ := reader.ReadString('\n') // Fixed input reading
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
