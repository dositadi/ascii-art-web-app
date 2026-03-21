package services

import (
	"net/http"
	"text/template"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderHomePage(w http.ResponseWriter, r *http.Request) *m.Error {
	temp, err := template.New("home.html").ParseFiles("web/static/internal_pages/home.html", "web/templates/home_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	homePageDetails := struct {
		AsciiPageUrl                                            string
		HeroBody                                                string
		SignUpUrl, LoginUrl                                     string
		IntroP1                                                 string
		IntroP2                                                 string
		IntroP3                                                 string
		LeadContributorName, Contributor2Name, Contributor3Name string
		LeadAbout, Contributor2About, Contributor3About         string
	}{
		AsciiPageUrl: h.ASCII_ROUTE,
		HeroBody: `ASCII (American Standard Code for Information Interchange) has been the backbone
      of digital communication since 1963. ASCIIForge lets you harness that legacy —
      transforming your words into stunning block art with a single click.
		`,
		SignUpUrl: h.SIGNUP_ROUTE, LoginUrl: h.LOGIN_ROUTE,
		IntroP1: `ASCII was created in 1963 as a standard character-encoding scheme for
        electronic communication. Using just 128 characters — letters, digits,
        punctuation and control codes — it became the universal language of
        computers worldwide.`,
		IntroP2: `ASCII art emerged in the 1960s and 70s as programmers discovered they
        could assemble characters into pictures, banners, and expressive
        typographic designs. Today it lives everywhere: terminal outputs,
        code comments, Discord servers, and developer tooling.`,
		IntroP3: `ASCIIForge brings that tradition into the modern web — giving you
        six distinct font styles, live preview, and one-click clipboard export.`,
		LeadContributorName: "Ositadinma Divine .A.", Contributor2Name: "Esther Lari", Contributor3Name: "Olayiwola Salawu",
		LeadAbout: `Divine built the authentication layer, session management, and the
              history-persistence API that backs every saved banner. He championed the
              zero-JS philosophy on the frontend and designed the RESTful API structure
              that makes ASCIIForge fast and stateless. Outside of work he contributes
              to open-source security tooling and mentors junior developers in his city.`,
		Contributor2About: `Esther led the visual identity of ASCIIForge — from the milk-and-blue palette
              to the monospace-driven component language. She drew the wireframes, iterated
              the card layouts, and obsessed over every 0.5px border and spacing token
              until the UI felt truly native to the terminal world. She's passionate about
              design systems and the intersection of print typography and web UI.`,
		Contributor3About: `Olayiwola is a full-stack developer with a deep love for the terminal aesthetic.
              He architected the ASCIIForge rendering engine, built the font style system
              from scratch, and designed the stateless CSS-only interaction model that
              powers the entire app with zero JavaScript. When he's not pushing pixels,
              he's writing about web performance and open-source tooling.`,
	}

	if s.GetHxRequestStatus(r) {
		if err2 := temp.ExecuteTemplate(w, "home", homePageDetails); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err2.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	} else {
		if err3 := temp.Execute(w, homePageDetails); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err3.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	}
	return nil
}
